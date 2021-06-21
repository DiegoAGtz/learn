package main

import "fmt"

const (
    a = 30
    b = 14.2
    c = true
    d = "AG"
    e int = 13
    f float64 = 93.8
    g bool = false
    h string = "DiegoAGtz"
)

func main() {
    fmt.Printf("%T\t%T\t%T\t%T\n", a, b, c, d)
    fmt.Printf("%T\t%T\t%T\t%T\n", e, f, g, h)
}
