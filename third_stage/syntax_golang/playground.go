package main

import "fmt"

var calTable = []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

func main() {
	c := make([]int, len(calTable))
	copy(c, calTable)

	for i := 0; i < len(calTable); i++ {
		fmt.Println("Value on ", i, " sm is: ", calTable[i], " ml")
	}
	for j := 0; j < len(c); j++ {
		if calTable[j]%20 == 0 {
			calTable = append(calTable, calTable[j]*2)
		}
	}
	fmt.Println(calTable)
}
