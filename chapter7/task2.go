package main

import "fmt"

func half(x int) (int, bool)  {
  halfx := x / 2
  return halfx, halfx % 2 == 0
}

func main()  {
  x := 3
  fmt.Println(half(x))
}
