package main

import "fmt"

func main() {
	// Массив нумерованная последовательность элементов одного типа с фиксированной длиной
	var a [5]int
	fmt.Println(a) // Массив с 5ю элементами типа int

	a[4] = 100
	fmt.Println(a)
	fmt.Println(a[4])

	fmt.Println("len: ", len(a)) // Длина массива

	b := [5]int{1, 2, 3, 4, 5} // Одновременное обьявление и инициализация массива
	fmt.Println("dcl:", b)

	var twoArrays [2][3]int // В одном многомерном массиве, два массива по три значения типа int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoArrays[i][j] = i + j
		}
	}
	fmt.Println("2 range array: ", twoArrays)

}
