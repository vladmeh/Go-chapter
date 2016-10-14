package main

import "fmt"

func main() {
  fmt.Print("Что будем конвертировать: 1 - температуру из градусов Фаренгейта в градусы Цельсия или 2 - футы в метры :")

  var i int
  fmt.Scanf("%v", &i)

  switch i {
    case 1: farToCel()
    case 2: futToMetr()
    default: fmt.Println("Unknown Number")
  }
}

func farToCel() {
  fmt.Print("Введите значение температуры по Фаренгейту: ")
  var f float64
  fmt.Scanf("%f", &f)
  c := (f - 32) * 5/9
  fmt.Println("Температура по Цельсию будет: ",c)
}

func futToMetr() {
  fmt.Print("Введите значение в футах: ")
  var input float64
  fmt.Scanf("%f", &input)
  m := input * .3048
  fmt.Println("В метах это будет: ",m)
}
