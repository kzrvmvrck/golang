package main

import "fmt"

func modifityArrays(arr [3]int) []int {
	arr[0] = 10
	return arr[:]
}

func arrSum(arr []int) int {
	result := 0
	for _, value := range arr {
		result += value
	}
	return result
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

	for index, value := range myArrAuto {
		fmt.Printf("value on index %d is: %d \n", index, value)
	}

	var x [5]int
	fmt.Println(x)

	x[3] = 666
	fmt.Println(x, len(x))
	fmt.Println(len(x))

	var xCode [58]string
	fmt.Println(xCode, len(xCode))

	for i := 65; i <= 122; i++ { //	65 индекс в таблице UTF это буква A
		xCode[i-65] = string(i)
	}

	fmt.Println(xCode)
	fmt.Println(xCode[42])

	// Опа ча! в Golang нет методов для  массивов, но можно определять функции для них

	// 1 Этап делаем функцию которая примет в аргументы Наш Массив. Вне функции main
	// arrSum

	// 2 Этап создаем массив
	testArr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	result := arrSum(testArr[:])
	fmt.Println(result)

	rvsk_1000_number_14 := []int{}
	fmt.Println(rvsk_1000_number_14)

	//Сравнение массивов
	// Массивы сравнимые сущности

}
