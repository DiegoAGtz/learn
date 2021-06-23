/*
    Crea un mapa con una llave tipo string el cual representa el 
    "nombre_apellido" de una persona y un vlaor de tipo []string el cual
    almacena sus cosas favoritas. Almacena tres registros en tu mapa. 
    Imprime todos sus valores con su índice de posición en el slice
    "eduar_tua", "computadoras", "montaña", "playa"
    "carlos_ramon", "leer", "comprar", "musica"
    "juan_bimba", "helado", "pintar", "bailar"
*/
package main

import "fmt"

func main() {
    x := map[string][]string {
        "eduar_tura"    : {"computadoras", "montaña", "playa"},
        "carlos_ramon"  : {"leer", "comprar", "musica"},
        "juan_bimba"    : {"helado", "pintar", "bailar"},
    }

    for i, v1 := range x {
        fmt.Println("Llave:", i)
        for j, v2 := range v1 {
            fmt.Printf("\tIndice: %v\tValor: %v\n", j, v2)
        }
    }
}
