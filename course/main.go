package main

import (
	"fmt"
)

const (
	A = iota
	B
	C
)
const ConstTemp = 5

func main() {
	//	printTypeDemo()

	//	callDivide()

	//printLevelDemo()

	//printConvertDemo()

	//printTypeSummary()

	//printSliceDemo()

	printArrayAndSliceDemo()
}

func printArrayAndSliceDemo() {
	arr := [3]string{"Tom", "Jack", "Lucy"}
	sli := []string{"Tom", "Jack", "Lucy"}

	fmt.Println("array:", arr)
	fmt.Println("slice:", sli)
	fmt.Println("切片比数组更常用，因为长度更灵活")
}
func printSliceDemo() {
	names := []string{"Tom", "Jack", "Lucy"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
}
func printTypeSummary() {
	fmt.Println("== 基础类型 ==")
	fmt.Println("int:", 1)
	fmt.Println("float64:", 3.14)
	fmt.Println("string:", "hello")
	fmt.Println("bool:", true)
	fmt.Println("== 除法差异 ==")
	fmt.Println("5/2 =", 5/2)
	fmt.Println("5.0/2.0 =", 5.0/2.0)
	fmt.Println("== 常量 iota ==")
	fmt.Println("A=", A, "B=", B, "C=", C)
	fmt.Println("== 类型转换 ==")
	fmt.Println("int to float64:", float64(1))
	fmt.Println("float64 to int:", int(1.0))

}
func printLevelDemo() {
	fmt.Println("A=", A, "B=", B, "C=", C)
}
func printConvertDemo() {
	x := 3
	y := 3.8
	fmt.Printf("x=%.1f\n", float64(x))
	fmt.Printf("y=%d\n", int(y))
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
