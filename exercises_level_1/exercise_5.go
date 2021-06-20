package main

import "fmt"

type second int

var x second
var y int

func main() {
    fmt.Printf("El valor de x es: %v, El tipo de x es: %T\n", x, x)
    x = 42
    fmt.Println(x)

    y = int(x)
    fmt.Printf("El valor de y es: %v, El tipo de y es: %T\n", y, y)
}
