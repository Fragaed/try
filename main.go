package main

import (
	"fmt"
	"strconv"
)

func main() {
	var aStr, bStr, operator string
	fmt.Println("Введите выражение (2 + 2):")
	fmt.Scan(&aStr, &operator, &bStr)

	a, _ := strconv.Atoi(aStr)
	b, _ := strconv.Atoi(bStr)

	if a <= 10 && a > 0 && b <= 10 && b > 0 {
		switch operator {
		case "+":
			println(a + b)
		case "-":
			println(a - b)
		case "*":
			println(a * b)
		case "/":
			println(a / b)
		}
	} else {
		println("error")
	}

}
