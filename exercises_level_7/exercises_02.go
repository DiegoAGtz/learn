/*
    Crea un struc persona
    Crea una función llamada "cambiame" con un *persona como
    parámetro.
    Cambia el valor almacenado en la dirección de memoria del *persona
        Para desreferenciar un struct se usa (*valor).campo
        Si el tipo de x es un tipo puntero con nombre y (*x).c es una expresión
        valida de seleción denotando un campo (pero no un método), x.c es una
        forma abreviada de (*x).c
    Crea un valor de tipo persona e imprimelo
    Llama a cambiame
    Vuelve a imprimir a persona
*/
package main

import "fmt"

type persona struct {
    nombre string
}

func main() {
    p1 := persona {
        nombre: "Ana",
    }
    fmt.Println(p1)
    cambiame(&p1)
    fmt.Println(p1)
}

func cambiame(p *persona) {
    (*p).nombre = "Anita"
    fmt.Println("Cambio interno:", *p)
    // Funciona con campos, no con métodos
    p.nombre = "Ana"
}
