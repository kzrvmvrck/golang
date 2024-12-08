package main

import "fmt"

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

}
