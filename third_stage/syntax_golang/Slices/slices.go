package main

import "fmt"

func appendSlice(baseSlice []string, value string) []string {
	newSlice := []string{}
	fmt.Println(newSlice)
	fmt.Println(len(baseSlice))

	for indexBaseSlice := 0; indexBaseSlice < len(baseSlice); indexBaseSlice++ {
		if indexBaseSlice == (len(baseSlice) - 1) {
			newSlice[len(baseSlice)-1] = value
		} else {
			newSlice[indexBaseSlice] = baseSlice[indexBaseSlice]
		}

	}
	return baseSlice
}

func main() {
	s := make([]string, 3) // Создание слисеров осуществляется при помощи функции make([]string, len). Указываются при это тип данных и дрина среза
	fmt.Println("emp slice:", s)

	s[0] = "a" // Так как срезы сущность индексируемая, мы можем присваивать значения элементов через индекс
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2]) // Может обратиться к элементу под конкретным индексом

	fmt.Println(len(s)) // Узнаем длину

	s = append(s, "d") // Срезы в отличии от массивов могут изменять свою длину
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s)) //Можно копировать один сред в другой
	copy(c, s)                  // сначала создаем пустой срезЮ потом в него закидываем значения

	fmt.Println("cpy: ", c)

	l1 := s[2:5]
	fmt.Println("sl1: ", l1)

	l2 := s[:5]
	fmt.Println("sl2: ", l2)

	l3 := s[2:]
	fmt.Println("sl3: ", l3)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl: ", t)

	twoSlices := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoSlices[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoSlices[i][j] = i + j
		}
	}
	fmt.Println("2sl: ", twoSlices)

	wer := make([]int, 5)

	fmt.Println(wer)

	arr := [5]int{1, 2, 3, 4, 5}
	slice := arr[:]

	fmt.Println(slice)

	slice = append(slice, 5)
	fmt.Printf("%d after used function append \n", slice)

	appArr := []string{"Jack", "Bryan", "Maks", "Tom", "Sergey", "Oleg"}
	fmt.Println(appArr)

	selfAppendArr := append(appArr, "Kyrt")
	fmt.Println(selfAppendArr)

	massAppendArr := append(appArr, "Devi", "Jons", "Alex")
	fmt.Println(massAppendArr)

	// Тут получается фигня. Он не меняет слайс базовый, в по сути у меня 3 разных

	// Пробуем через функцию
	appendSlice(appArr, "Kyrt")
	fmt.Println(appArr)

}
