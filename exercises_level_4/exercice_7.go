/*
    Crear un slice de slice de string. Almacena los siguientes valores
    en un slice multi-dimensional:
        "Batman", "Jefe", "Vestido de negro"
        "Robin", "Ayudante", "Vestido de colores"
    Haz range sobre ellos, luego sobre los datos de cada elemento
*/
package main

import "fmt"

func main() {
    gotham := [][]string { {"Batman", "Jefe", "Vestido de negro"}, {"Robin", "Ayudante", "Vestido de colores"} }

    for i, v1 := range gotham {
        fmt.Println("Slice numero: ", i)
        for _, v2 := range v1 {
            fmt.Printf("\t\t %v\n", v2)
        }
    }
}
