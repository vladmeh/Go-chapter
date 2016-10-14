package main

import "fmt"

func main() {
  x := []int{
    48, 96, 86, 68,
    57, 82, 63, 70,
    37, 34, 83, 27,
    19, 97, 9, 17,
  }

  max := 0
  // for i := 0; i < len(x); i++ {
  //   if x[i] > max {
  //     max = x[i]
  //   }
  // }
  for _, value := range x {
    if value > max {
      max = value
    }
  }
  fmt.Println("Максимальное число в массиве:", max)

  min := max
  // for i := 0; i < len(x); i++ {
  //   if x[i] < min {
  //     min = x[i]
  //   }
  // }
  for _, value := range x {
    if value < min {
      min = value
    }
  }
  fmt.Println("Минимальное число в массиве:", min)
}
