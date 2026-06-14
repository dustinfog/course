package main

import "fmt"

func main() {
	printStudentDemo()

}

type Student struct {
	Name   string
	Scores []int
}

func printStudentDemo() {
	stu1 := Student{
		Name:   "Tom",
		Scores: []int{90, 80, 100},
	}
	stu2 := Student{
		Name:   "Lucy",
		Scores: []int{90, 80, 100},
	}

	fmt.Println("name:", stu1.Name)
	fmt.Println("scores:", stu1.Scores)
	fmt.Println("name:", stu2.Name)
	fmt.Println("scores:", stu2.Scores)
}
