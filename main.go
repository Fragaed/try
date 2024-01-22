package main

import (
	"fmt"
	"math"
	"strconv"
)

var aStr, bStr, operator string

func main() {
	var name string
	var function int
	var i string
	fmt.Println("What is your name?")
	fmt.Scanln(&name)
	fmt.Println("Hello", name, "change function: \n 1. Калькулятор \n 2. Нахождение корней квадратного уровнения")
	fmt.Scanln(&function)
	switch {
	case function == 1:
		calk()
	case function == 2:
		poiskx()
	}
	fmt.Println("Желаете продолжить? (y/n)")
	fmt.Scanln(&i)

	for i == "y" {
		fmt.Println("Hello", name, "change function: \n 1. Калькулятор \n 2. Нахождение корней квадратного уровнения")
		fmt.Scanln(&function)
		switch {
		case function == 1:
			calk()
		case function == 2:
			poiskx()
		}
		fmt.Println("Желаете продолжить? (y/n)")
		fmt.Scanln(&i)

	}
	fmt.Scanln()

}

func calk() {
	fmt.Println("Введите выражение (2 + 2) (считает числа от 1 до 10:")
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
	fmt.Scanln()
}

func poiskx() {
	var a, b, c float64

	fmt.Println("    Решаем квадратное уровнение")
	fmt.Println("Введи a:")
	fmt.Scan(&a)
	fmt.Println("Введи b:")
	fmt.Scan(&b)
	fmt.Println("Введи c:")
	fmt.Scan(&c)

	D := (b * b) - 4*(a*c)

	if D > 0 {
		var x1 float64
		var x2 float64

		x1 = (-b + math.Sqrt(D)/(2*a))
		x2 = (-b - math.Sqrt(D)/(2*a))

		fmt.Println("Ваше уровнение имеет 2 корня \nD = " + fmt.Sprint(D))
		fmt.Println("X1: " + fmt.Sprint(x1) + "\nX2: " + fmt.Sprint(x2))

	} else if D == 0 {
		var x float64

		x = (-b) / (2 * a)

		fmt.Println("Ваше уровнение имеет 1 корень \n D=0")
		fmt.Println("x: " + fmt.Sprint(x))
	} else if D < 0 {
		fmt.Println("Ваше уровнение не имеет корней \nD: " + fmt.Sprint(D))
	}
	fmt.Scanln()
}
