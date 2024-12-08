package main

import "fmt"

func main() {

    //Переменные в golang
    var a string = "initial"   // Строковая переменная
    fmt.Println(a)

    var b, c int = 1, 2
    fmt.Println(b, c)           // Числовые переменные

    var trust = true
    fmt.Println(trust)          // Булевые значения

    var e int
    fmt.Println(e)              // Если переменная объявлена, но не инициилизирована, то ей присваивается нулевое
                                // значение. В случае и int это будет 0

    var eStr string
    fmt.Println(eStr)           // В консоле ничего, потом стоит разобраться почему!

    f := "short"
    fmt.Println(f)              // := это сокращение для объявления и инициализации переменной (vat f string = "short")

    // Из иного источника
    // Переменные возможно объявлять блочным способом
    var (
        age     int     =   50
        name    string  =   "Jack"
        height  float32 =   171.2
        weight  float32 =   70.2
        married bool    =   false
    )

    fmt.Println(age, name, height, weight, married)


}