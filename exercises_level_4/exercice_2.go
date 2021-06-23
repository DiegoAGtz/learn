/*
    Usando un composite literal
        Crea un slice de tipo int 
        Asigna 10 valores
    Itera con range sobre el slice e imprime los valores
    Imprime el tipo del arreglo
*/
package main

import "fmt"

func main() {
    x := []int { 0, 9, 7, 2, 1, 10, 99, 921, 3, 4 }
    for i, v := range x {
        fmt.Printf("Indice: %v - Valor: %v\n", i, v)
    }
    fmt.Printf("Tipo de variable: %T\n", x)
}
