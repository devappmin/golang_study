package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)

	people := [2]string{"kim", "lee"}

	for _, person := range people {
		go isAwesome(person, c)
	}

	for i := 0; i < len(people); i++ {
		fmt.Println("Waiting for messages:", i)
		result := <-c
		fmt.Println("Received this message:", result)
	}

}

func isAwesome(person string, c chan string) {
	time.Sleep(time.Second * 5)
	c <- person + " is awesome!"
}
