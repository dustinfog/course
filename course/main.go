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

	//printArrayAndSliceDemo()

	//printRangeDemo()

	//printAppendDemo()

	//printNumberSlice()

	numss := []int{10, 20, 30, 40}
	result := sumSlice(numss)
	fmt.Println("sum of numss:", result)
}

func sumSlice(nums []int) int {
	sum := 0
	for _, value := range nums {
		sum += value
	}
	return sum
}

func printNumberSlice() {
	nums := []int{10, 20, 30, 40}
	sum := 0
	for _, value := range nums {
		sum = sum + value
	}
	fmt.Println("nums:", nums)
	fmt.Println("sum:", sum)

	nums1 := []int{1, 2, 3, 4, 5}
	sums1 := 0
	for _, value := range nums1 {
		sums1 += value
	}
	fmt.Println("nums1:", nums1)
	fmt.Println("sums1:", sums1)
}

func printAppendDemo() {
	names := []string{"Tom", "Jack", "Lucy"}
	fmt.Println("before:", names)

	names = append(names, "Mike")
	names = append(names, "Lily")

	fmt.Println("after:", names)

	names = append(names, "Ben")
	fmt.Println("len:", len(names))
}

func printRangeDemo() {
	names := []string{"Tom", "Jack", "Lucy"}

	for index, value := range names {
		fmt.Println(index, value)
	}

	nums := []int{10, 20, 30}

	for index, value := range nums {
		fmt.Println(index, value)
	}
}

func printSliceDemo() {
	names := []string{"Tom", "Jack", "Lucy"}
	for index := 0; index < len(names); index++ {
		value := names[index]

		fmt.Println(index, value)
	}
}
func printArrayAndSliceDemo() {
	arr := [3]string{"Tom", "Jack", "Lucy"}
	sli := []string{"Tom", "Jack", "Lucy"}

	fmt.Println("array:", arr)
	fmt.Println("slice:", sli)
	fmt.Println("切片比数组更常用，因为长度更灵活")
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
