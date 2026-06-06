package main

import (
	"fmt"
)

const ConstTemp = 5

func main() {
	printTypeDemo()

	callDivide()
}

func printTypeDemo() {
	var a int = 1
	var b float64 = 2
	var c string = "hello"
	var d bool = true

	fmt.Println("a=", a)
	fmt.Printf("a=%d\n", a)
	fmt.Println("b=", b)
	fmt.Printf("b=%.2f\n", b)
	fmt.Printf("c=%q\n", c)
	fmt.Printf("d=%t\n", d)
	fmt.Printf("5/2=%d\n", 5/2)
	fmt.Printf("5.0/2.0=%.2f\n", 5.0/2.0)
}

func callSafeModulo() {
	r3, err3 := safeModulo(100, ConstTemp)
	if err3 != nil {
		fmt.Println("Error:", err3)
	} else {
		fmt.Println("result:", r3)
	}
	r4, err4 := safeModulo(100, 0)
	if err4 != nil {
		fmt.Println("Error:", err4)
	} else {
		fmt.Println("result:", r4)
	}
}

func safeModulo(base, i int) (int, error) {
	if i == 0 {
		return 0, fmt.Errorf("i不能是0")
	}
	return base % i, nil
}

func callDivide() {
	r1, err1 := divide(100, ConstTemp)
	if err1 != nil {
		fmt.Println("error:", err1)
	} else {
		fmt.Println("result:", r1)
	}

	r2, err2 := divide(100, 0)
	if err2 != nil {
		fmt.Println("error:", err2)
	} else {
		fmt.Println("result:", r2)
	}
}

func divide(base int, i int) (int, error) {
	if i == 0 {
		return 0, fmt.Errorf("i不能是0")
	}
	return base / i, nil
}
