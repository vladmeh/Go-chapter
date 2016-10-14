package main

import "fmt"

func main() {
  fmt.Print("Введите значение температуры по Фаренгейту: ")
  var f float64
  fmt.Scanf("%f", &f)

  c := (f - 32) * 5/9

  fmt.Println("Температура по Цельсию будет: ",c)
}
