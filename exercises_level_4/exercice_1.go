/*
    Usando un composite literal
        Crea un arreglo de tipo int el cual tenga 5 valores
    Itera con range sobre el arreglo e imprime los valores
    Imprime el tipo del arreglo
*/
package main

import "fmt"

func main() {
    x := [5]int { 0, 9, 7, 2, 1 }
    for i, v := range x {
        fmt.Printf("Indice: %v - Valor: %v\n", i, v)
    }
    fmt.Printf("Tipo de variable: %T\n", x)
}
