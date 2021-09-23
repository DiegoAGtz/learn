/*
    Métodos
*/

package main

import "fmt"

type persona struct {
    nombre      string
    apellido    string
    edad        int
}

func main() {
    p1 := persona {
        nombre: "Diego Armando",
        apellido: "Gutiérrez Ayala",
        edad: 21,
    }

    p2 := persona {
        nombre: "Anastasia",
        apellido: "Steele",
        edad: 20,
    }
    p1.presentar()
    p2.presentar()
}

func (p persona) presentar() {
    fmt.Println("Hola, mi nombre es:", p.nombre, p.apellido, "y tengo:", p.edad, "años")
}
