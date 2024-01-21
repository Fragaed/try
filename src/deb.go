package main

import (
	"fmt"
)

func main() {
	fmt.Println("			Дебильный калькулятор")

	var action string
	var a float64
	var b float64

	fmt.Println("Введите первое число: ")
	fmt.Scan(&a)

	fmt.Println("Какое действие вы хотите выполнить? (+, -, *, /, )")
	fmt.Scan(&action)

	fmt.Println("Введите второе число: ")
	fmt.Scan(&b)

	switch {
	case action == "+":
		fmt.Println("Сумма чисел =" + fmt.Sprint(a+b))
	case action == "-":
		fmt.Println("Разность чисел =" + fmt.Sprint(a-b))
	case action == "*":
		fmt.Println("Результат умножения чисел =" + fmt.Sprint(a*b))
	case action == "/":
		fmt.Println("Результат диления чисел =" + fmt.Sprint(a/b))
	}
	var x int
	fmt.Scan(&x)
}
