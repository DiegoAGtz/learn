/*
    Funcionamiento de defer.
    Ejecuta la función con los argumentos que se le pasan al momento
    de llamarla.
*/
package main

import "fmt"

func main() {
    a, b := 42, 10
    defer foo(a, b)
    
    a = 25
    b = 15

    foo(a, b)
    fmt.Println("Hola mundo desde la función main()")
}

func foo(a, b int) {
    defer func() {
        fmt.Println("Hola desde la función anónima")
    }()
    fmt.Println("Hola mundo desde la función foo()", a, b)
}
