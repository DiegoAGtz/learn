/*
    Para borrar elementos de un slice, usamos append en conjunto
    con slicing (dividir). Para este ejercicio, sigue los siguientes pasos

    Comienza con un slice
        x := []int { 42, 43, 44, 45, 46, 47, 48, 49, 50, 51 }
    Usa append y slicing para obtener los siguientes valores, los cuales, 
    debes asignar a una variable "y" y luego imprimir.
        { 42, 43, 44, 48, 49, 50, 51 }
*/
package main

import "fmt"

func main() {
    x := []int { 42, 43, 44, 45, 46, 47, 48, 49, 50, 51 }
    y := append(x[:3] , x[6:]...)
    
    fmt.Println(y)
}
