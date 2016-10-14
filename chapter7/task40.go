package main

import "fmt"

func factorial(x uint) uint {
    if x == 0 {
        return 1
    }

    return x * factorial(x-1)
}

func main() {
    fmt.Print("Введите число: ")
    var x uint
    fmt.Scanf("%v", &x)

    fmt.Println("Факториал числа", x,"равен:", factorial(x))
}
