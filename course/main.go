package main

import (
	"fmt"
)

func main() {

	r1, err1 := divide(100, 5)
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
