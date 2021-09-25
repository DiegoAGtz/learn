/*
    Crea una variable de cualquier tipo e imprime su dirección
    de memoria
*/
package main

import "fmt"

func main() {
    a := 42
    fmt.Println(&a)
}
