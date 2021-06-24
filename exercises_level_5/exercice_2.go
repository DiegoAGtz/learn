/*
    Usa el código del ejemplo anterior y almacena los valores de tipo persona en
    un mapa con la llave apellido. Accede a cada valor en el mapa. Imprime los valores.
    Imprime también los valores haciendo range sobre el slice.
*/
package main

import "fmt"

type persona struct {
    nombre string
    apellido string
    saboresFav []string
}

func main() {
    p1 := persona {
        nombre: "Diego",
        apellido: "Gutierrez",
        // Usamos composite literal
        saboresFav: []string{
            "Fresa", 
            "Vainilla", 
            "Limon", 
            "Galleta",
        },
    }
    p2 := persona {
        nombre: "Armando",
        apellido: "Ayala",
        saboresFav: []string{
            "Galleta",
            "Chocolate", 
            "Coco", 
            "Chicle", 
            "Frutas",
        },
    }

    m := map[string]persona {
        p1.apellido: p1,
        p2.apellido: p2,
    }

    for _, item := range m {
        fmt.Println(item.nombre, item.apellido)
        for _, v := range item.saboresFav {
            fmt.Println("\t", v)
        }
        fmt.Println("-----------------")
    }
    
}
