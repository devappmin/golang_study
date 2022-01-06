package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	favFood := []string{"sushi", "kimbap", "ramen", "rice", "eggplant"}
	kim := person{
		name:    "kim seung hwan",
		age:     26,
		favFood: favFood,
	}
	fmt.Println(kim)
}
