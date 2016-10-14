package main

import ("fmt"; "math")

type Circle struct {
    x, y, r float64
}

func distance(x1, y1, x2, y2 float64) float64 {
    a := x2 - x1
    b := y2 - y1
    return math.Sqrt(a*a + b*b)
}
func rectangleArea(x1, y1, x2, y2 float64) float64 {
    l := distance(x1, y1, x1, y2)
    w := distance(x1, y1, x2, y1)
    return l * w
}

func (c *Circle) area() float64 {
    return math.Pi * c.r*c.r
}

func (c *Circle) length() float64 {
    return math.Pi*2 * c.r
}

func main() {
    var rx1, ry1 float64 = 0, 0
    var rx2, ry2 float64 = 10, 10
    fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))

    c := Circle{0, 0, 5}
    fmt.Println("Площадь круга:", c.area())
    fmt.Println("Длина окружности:", c.length())
}
