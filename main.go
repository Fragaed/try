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
	fmt.Println("Как тебя зовут?")
	fmt.Scanln(&name)
	fmt.Println("Привет", name, "выбери функцию: \n 1. Калькулятор \n 2. Нахождение корней квадратного уровнения")
	fmt.Println(" 3. Подсчет оченки за четверть по текущим.")
	fmt.Scanln(&function)
	switch {
	case function == 1:
		calk()
	case function == 2:
		poiskx()
	case function == 3:
		ocenka()
	}
	fmt.Println("Желаете продолжить? (д/н)")
	fmt.Scanln(&i)

	for i == "д" {
		fmt.Println("Привет", name, "выбери функцию: \n 1. Калькулятор. \n 2. Нахождение корней квадратного уровнения.")
		fmt.Println(" 3. Подсчет оченки за четверть по текущим.")
		fmt.Scanln(&function)
		switch {
		case function == 1:
			calk()
		case function == 2:
			poiskx()
		case function == 3:
			ocenka()
		}
		fmt.Println("Желаете продолжить? (д/н)")
		fmt.Scanln(&i)

	}
	fmt.Scanln()

}

func calk() {
	fmt.Println("Введите выражение (2 + 2) (считает числа от 1 до 10):")
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

func ocenka() {
	var x float64
	sum := 0.0
	var srednee float64
	marks := []float64{}
	fmt.Println("Введите ваши оценки: (Чтобы прекратитить введите число >5)")
	for i := 0; x <= 5; i++ {
		fmt.Scan(&x)
		marks = append(marks, x)
		sum = sum + marks[i]
	}
	sum = sum - marks[len(marks)-1]
	marks = marks[:len(marks)-1]
	kol_vo := float64(len(marks))
	srednee = math.Round(sum / kol_vo)
	fmt.Println("Список оценок: \n", marks)
	fmt.Println("Колличество оценок: ", len(marks))
	fmt.Println("Оценка которая выходит в четверти: ", srednee)
	fmt.Scanln()
}
