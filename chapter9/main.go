package main

import ("fmt"; "math")

type Circle struct {
    x, y, r float64
}

type Rectangle struct {
    x1, y1, x2, y2 float64
}

type Person struct {
    Name string
}

type Android struct {
    Person
    Model string
}

type Shape interface {
    area() float64
    perim() float64
}

type MultiShape struct {
	shapes []Shape
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

func (r *Rectangle) perim() float64 {
    l := distance(r.x1, r.y1, r.x1, r.y2)
    w := distance(r.x1, r.y1, r.x2, r.y1)
    return 2*l + 2*w
}

func (c *Circle) area() float64 {
    return math.Pi * c.r*c.r
}

func (c *Circle) perim() float64 {
    return math.Pi*2 * c.r
}

func (p *Person) Talk() {
    fmt.Println("Hi, my name is", p.Name)
}

func totalCalc(shapes ...Shape) (float64, float64) {
    var area float64
    var perim float64
    for _, s := range shapes {
        area += s.area()
        perim += s.perim()
    }
    return area, perim
}

func (m *MultiShape) calc() (float64, float64) {
	var area float64
	var perim float64
	for _, s := range m.shapes {
		area += s.area()
		perim += s.perim()
	}
	return area, perim
}

func main() {
    r := Rectangle{0, 0, 10, 5}
    fmt.Println("Площадь прямоугольника:", r.area())
    fmt.Println("Периметр прямоугольника:", r.perim())

    fmt.Println("------")
    c := Circle{0, 0, 5}
    fmt.Println("Площадь круга:", c.area())
    fmt.Println("Перимерт окружности:", c.perim())

    fmt.Println("---totalCalc---")
    tarea, tperim := totalCalc(&c, &r)
    fmt.Println("Общая площадь всех фигур:", tarea)
    fmt.Println("Oбщий периметр всех фигур:", tperim)

    fmt.Println("---MultiShape---")
    mshape := new(MultiShape)
    mshape.shapes = []Shape{&c, &r}
    marea, mperim := mshape.calc()
    fmt.Println("Общая площадь всех фигур:", marea)
    fmt.Println("Oбщий периметр всех фигур:", mperim)

    fmt.Println("------")
    n := Person{"Вася"}
    //a := new(Android)
    a := Android{n, "Пупкин"}
    fmt.Println(a.Name, a.Model)
    a.Talk()
}
