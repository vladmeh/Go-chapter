package main

import "fmt"
import "github.com/golang-book/chapter11/math"

func main() {
    xs := []float64{1,2,3,4,5}
    avg := math.Average(xs)
    max := math.Max(xs)
    min := math.Min(xs)
    fmt.Println("Cреднее в массиве чисел:", avg)
    fmt.Println("Максимальное число в массиве:", max)
    fmt.Println("Минимальное число в массиве:", min)
}
