/*
    Crea tu propio tipo persona el cual tendrá un tipo subyacente struct de manera
    que pueda almacenar la siguiente data:
        Nombre
        Apellido
        Sabores de helado favoritos
    Crea dos vaclores de tipo persona. Imprime los valores, usa range sobre los sabores
    de helado.
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

    fmt.Println(p1.nombre, p1.apellido)
    for _, v := range p1.saboresFav {
        fmt.Printf("\t%v\n", v)
    }
    fmt.Println(p2.nombre, p2.apellido)
    for _, v := range p2.saboresFav {
        fmt.Printf("\t%v\n", v)
    }
    
}
