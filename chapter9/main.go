package main

import ("fmt"; "math")

type Circle struct {
    x, y, r float64
}

type Rectangle struct {
    x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
    a := x2 - x1
    b := y2 - y1
    return math.Sqrt(a*a + b*b)
}
func (r *Rectangle) area() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return l * w
}

func (c *Circle) area() float64 {
    return math.Pi * c.r*c.r
}

func (c *Circle) length() float64 {
    return math.Pi*2 * c.r
}

func main() {
    r := Rectangle{0, 0, 10, 10}
    fmt.Println("Площадь прямоугольника:", r.area())

    c := Circle{0, 0, 5}
    fmt.Println("Площадь круга:", c.area())
    fmt.Println("Длина окружности:", c.length())
}
