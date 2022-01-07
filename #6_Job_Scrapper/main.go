package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	company  string
	salary   string
	summary  string
	upload   string
}

var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	jobs := []extractedJob{}
	c := make(chan []extractedJob)
	pages := getPages()

	for i := 0; i < pages; i++ {
		go getPage(i, c)
	}

	for i := 0; i < pages; i++ {
		extractedJobs := <-c
		jobs = append(jobs, extractedJobs...)
	}

	writeJobs(jobs)
	fmt.Println("Done, extracted:", len(jobs))
}

func writeJobs(jobs []extractedJob) {
	c := make(chan []string)

	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	header := []string{"Link", "title", "location", "company", "salary", "summary", "upload"}

	wErr := w.Write(header)
	checkErr(wErr)

	for _, job := range jobs {
		go writeJob(job, c)
	}

	for i := 0; i < len(jobs); i++ {
		jobSlice := <-c
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func writeJob(job extractedJob, c chan<- []string) {
	c <- []string{
		"https://kr.indeed.com/viewjob?jk=" + job.id,
		job.title,
		job.location,
		job.company,
		job.salary,
		job.summary,
		job.upload,
	}
}

func getPage(page int, mainC chan<- []extractedJob) {
	jobs := []extractedJob{}
	c := make(chan extractedJob)

	pageUrl := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageUrl)

	res, err := http.Get(pageUrl)

	checkErr(err)
	checkCode(res)

	// Close the connection after the function returns.
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".tapItem")

	searchCards.Each(func(i int, card *goquery.Selection) {

		go extractJob(card, c)
		// jobs = append(jobs, job)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-jk")
	title := cleanString(card.Find(".jobTitle>span").Text())
	company := cleanString(card.Find(".companyName").Text())
	location := cleanString(card.Find(".companyLocation").Text())
	salary := cleanString(card.Find(".salary-snippet").Text())
	summary := cleanString(card.Find(".job-snippet").Text())
	upload := cleanString(card.Find(".date").Text())

	c <- extractedJob{
		id:       id,
		title:    title,
		company:  company,
		location: location,
		salary:   salary,
		summary:  summary,
		upload:   upload,
	}
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)

	checkErr(err)
	checkCode(res)

	// Close the connection after the function returns.
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed:", res.Status)
	}
}
