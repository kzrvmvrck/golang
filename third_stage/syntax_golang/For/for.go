package main

import "fmt"

func main() {
	// For это единственная конструкция циклов в golang
	i := 1
	// Базовый тип цикла с единственным условие
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	// Классический вариант: значение/условие/постусловие
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}
	// Цикл без условия. Будет работать до тех пока, пока не будет применен break или return внутри функции
	var counter int = 0
	for {
		fmt.Println("Loop")
		if counter == 3 {
			break
		}
		counter++
	}

}
