/*
    Usando el código del ejercicio anterior, agrega un registro a tu mapa,
    ahora imprime el mapa usando range
*/
package main

import "fmt"

func main() {
    x := map[string][]string {
        "eduar_tura"    : {"computadoras", "montaña", "playa"},
        "carlos_ramon"  : {"leer", "comprar", "musica"},
        "juan_bimba"    : {"helado", "pintar", "bailar"},
    }

    x["AG"] = []string { "programar", "jugar", "leer" }

    for i, v1 := range x {
        fmt.Println("Llave:", i)
        for j, v2 := range v1 {
            fmt.Printf("\tIndice: %v\tValor: %v\n", j, v2)
        }
    }

}
