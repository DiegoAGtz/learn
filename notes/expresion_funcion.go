package main

import "fmt"

func main() {
    f := func() {
        fmt.Println("Mi primera expresión función")
    }
    f()

    g := func(x int) {
        fmt.Println("El año en que se descubrió América fué: ", x)
    }

    g(1492)
}

