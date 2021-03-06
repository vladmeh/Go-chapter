package main

import "fmt"
import "time"

func sleep(t int, done chan bool) {
    fmt.Print("ждем ", t, " секунд...")

    <- time.After(time.Second * time.Duration(t))
    fmt.Println("ок")

    done <- true
}

func main() {
    done := make(chan bool, 1)
    go sleep(5, done)
    <- done
}
