package main

import "fmt"

func sum(xs []int) int {
    total := 0
    for _, v := range xs {
        total += v
    }
    return total
}

func main() {
    xs := []int{1,2,3,4,5}
    fmt.Println(sum(xs))
}
