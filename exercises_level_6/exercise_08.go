/*
    Crear una función que retorne una función.
    Asignar la función retornada a una variable.
    Llama la función retornada.
*/

package main

import "fmt"

func main() {
    a := foo()

    fmt.Println(a())
    fmt.Println(a())
    fmt.Println(a())
}

func foo() func() int {
    x := 42
    return func() int {
        x++
        return x
    }
}
