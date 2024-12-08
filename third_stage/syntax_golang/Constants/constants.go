package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 50000000 // Оператор const может быть применет везде где можно применить оператор var

	const d = 3e20 / n // Арифметические действия с константами производятся с произвольной точностью
	fmt.Println(d)

	fmt.Println(int64(d)) // Чиловые константы не имеют типа до тех пор, пока он не будет определен явно

	fmt.Println(math.Sin(n)) // Числу может быть присвоен тип в контексте (в функции)

}
