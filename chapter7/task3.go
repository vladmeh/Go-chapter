package main

import "fmt"

func makeEvenGenerator() func() uint {
    i := uint(0)
    return func() (ret uint) {
        ret = i
        i += 2
        return
    }
}
func makeOddGenerator() func() uint {
    i := uint(11)
    return func() (ret uint) {
        ret = i
        i += 2
        return
    }
}
func main() {
    nextOdd := makeOddGenerator()
    fmt.Println(nextOdd()) // 0
    fmt.Println(nextOdd()) // 2
    fmt.Println(nextOdd()) // 4
}
