package main

import "fmt"

func printSeries(limit int) {
	for i := 1; i <= limit; i++ {
		fmt.Println("i =", 100/i)
	}
}

func main() {
	s := "gopher"
	fmt.Printf("Hello and welcome, %s!\n", s)

	printSeries(10)
}
