package main

import "fmt"

func modifityArrays(arr [3]int) []int {
	arr[0] = 10
	return arr[:]
}

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

	// Уроки chatagpt
	var myArr [5]int
	myArr[0] = 12
	myArr[1] = 13
	myArr[2] = 14
	myArr[3] = 15
	myArr[4] = 16

	fmt.Println(myArr)

	var myArrOne = [5]int{1, 2, 3, 4, 5}
	fmt.Println(myArrOne, len(myArrOne))

	xRes := len(myArrOne)
	xResOne := int(xRes) - 1

	fmt.Println(xResOne)

	// Определение массива автоматически
	myArrAuto := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(myArrAuto, len(myArrAuto))

	for index := 0; index < len(myArrAuto); index++ {
		fmt.Println(myArrAuto[index])
	}

	arrTestInFunc := [3]int{1, 2, 3}
	arrTestInFuncOne := modifityArrays(arrTestInFunc)
	fmt.Println(arrTestInFuncOne, arrTestInFunc)

}
