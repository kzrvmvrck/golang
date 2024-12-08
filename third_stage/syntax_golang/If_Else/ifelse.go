package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 Четное число")
	} else {
		fmt.Println("7 Нечетное число")
	}

	if 8%2 == 0 {
		fmt.Println("8 Четное число")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "Отрицательное число")
	} else if num < 10 {
		fmt.Println(num, "Число состоит из одной цифры")
	} else if num > 10 {
		fmt.Println(num, "Cостоит из дву чисел")
	}
}
