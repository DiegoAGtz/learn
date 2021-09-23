/*
    Interfaces.
    Crear un tipo cuadrado y un tipo circulo.
    Adjuntar un método a cada uno que calcule y retorne el área
    Crear un tipo forma que defina una interfaz como cualquier cosa que tenga
    el método área.
    Crear una función info que toma un tipo forma e imprime el área de la misma
*/

package main

import (
    "fmt"
    "math"
)

type cuadrado struct {
    lado float64
}

type circulo struct {
    radio float64
}

type forma interface {
    area() float64
}

func main() {
    c1 := cuadrado {
        lado: 5,
    }
    
    c2 := circulo {
        radio: 1,
    }

    info(c1)
    info(c2)
}

func (c cuadrado) area() float64 {
    return c.lado * c.lado
}

func (c circulo) area() float64 {
    return math.Pi * c.radio * c.radio
}

func info(f forma) {
    fmt.Println("El área es:", f.area())
}
