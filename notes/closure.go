/*
    Los closures nos ayudan a limitar el scope de las variables
    Es importante tener el scope de las variables reducido
*/

package main

import "fmt"

var x int

func main() {
    fmt.Println(x)
    x++
    fmt.Println(x)
    foo()
    fmt.Println(x)

    {
        /*
            Scope interno
        */
        y := 42
        fmt.Println(y)
    }

    
    a := incrementor()
    b := incrementor()
    fmt.Println("a:", a())
    fmt.Println("a:", a())
    fmt.Println("a:", a())
    fmt.Println("a:", a())
    fmt.Println("a:", a())
    fmt.Println("b:", b())
    fmt.Println("b:", b())
    fmt.Println("b:", b())
    fmt.Println("b:", b())

}

func foo() {
    fmt.Println("Hola")
    x++
}

func incrementor() func() int {
    var x int
    // Closure
    return func() int {
        x++
        return x
    }
}
