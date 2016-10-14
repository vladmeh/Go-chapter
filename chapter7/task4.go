package main

import "fmt"

func fib(n int) int {
    if n == 0 || n == 1 {
        return 1
    }

    return fib(n-1) + fib(n-2)
}

func main() {
    fmt.Print("Введите номер ряда Фиббоначи: ")
    var n int
    fmt.Scanf("%v", &n)

    fmt.Println("Число Фиббоначи в", n,"ряду равно:", fib(n))
}
