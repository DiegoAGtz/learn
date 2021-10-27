package main

import (
    "fmt"
    "sort"
)

type persona struct {
    nombre string
    edad int
}

type porEdad []persona
type porNombre []persona

// Para ordenar personalizadamente tenemos que usar las siguientes
// funciones que pertenecen a la interface Interface del paquete sort
func (a porEdad) Len() int {return len(a)}
func (a porEdad) Swap(i, j int) {a[i], a[j] = a[j], a[i]}
func (a porEdad) Less(i, j int) bool {return a[i].edad < a[j].edad}

func (a porNombre) Len() int {return len(a)}
func (a porNombre) Swap(i, j int) {a[i], a[j] = a[j], a[i]}
func (a porNombre) Less(i, j int) bool {return a[i].nombre < a[j].nombre}

func main() {
    s1 := []int{4, 24, 57, 2, 7, 73, 4, 64, 34, 1}
    s2 := []string{"Diego", "Go", "Ana", "Hola", "Como", "Miau"}
    
    fmt.Println(s1)
    sort.Ints(s1)
    fmt.Println(s1)

    fmt.Println(s2)
    sort.Strings(s2)
    fmt.Println(s2)

    p1 := persona {
        nombre: "Diego",
        edad: 20,
    }
    p2 := persona {
        nombre: "Ana",
        edad: 21,
    }
    p3 := persona {
        nombre: "Ani",
        edad: 19,
    }
    personas := []persona{p1, p2, p3}

    fmt.Println(personas)
    sort.Sort(porNombre(personas))
    fmt.Println(personas)

    sort.Sort(porEdad(personas))
    fmt.Println(personas)
}
