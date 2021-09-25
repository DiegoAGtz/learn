/*
   Closures
   Encerrar el scope de una variable en un bloque de código
*/

package main

import (
    "fmt"
)

func main() {
    f := foo()
    fmt.Println(f())
    fmt.Println(f())
    fmt.Println(f())
    fmt.Println(f())
    fmt.Println(f())
}

func foo() func() int {
    x := 0
    return func() int {
        x++
        return x
    }
}
