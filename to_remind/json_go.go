package main

import (
	"encoding/json"
	"fmt"
)

// Para que pueda ser convertido por Marshal necesitá tener mayusculas.
// Sigue el mismo principio de las funciones públicas
type persona struct {
    Nombre string
    Apellido string
    Edad int
}

func main() {
    fmt.Println("---------")
    fmt.Println("Marshal")
    p1 := persona {
        Nombre: "Diego Armando",
        Apellido: "Gutierrez Ayala",
        Edad: 20,
    }
    p2 := persona {
        Nombre: "Ana",
        Apellido: "Steel",
        Edad: 20,
    }
    personas := []persona{p1, p2}
    fmt.Println(personas)

    bs, err := json.Marshal(personas)
    if err != nil {
        fmt.Println("Error:", err)
    }
    fmt.Println(string(bs))
    fmt.Println(bs)

    fmt.Println("")
    fmt.Println("---------")
    fmt.Println("Unmarshal")

    valores := `[{"Nombre":"Diego Armando","Apellido":"Gutierrez Ayala","Edad":20},{"Nombre":"Ana","Apellido":"Steel","Edad":20}]`
    bs2 := []byte(valores)
    var personas2 []persona

    err = json.Unmarshal(bs2, &personas2)
    if err != nil {
        fmt.Println("Error:", err)
    }
    fmt.Println(personas2)
    for i, v := range personas2 {
        fmt.Println("Datos persona:", i)
        fmt.Println(v.Nombre, v.Apellido, v.Edad)
    }
}
