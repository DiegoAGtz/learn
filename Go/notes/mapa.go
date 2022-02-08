package main

import "fmt"

func main() {
    // Maps can be used to data search
    m := map[string]int {
        "Batman": 32,
        "Robin" : 27,
        "AG"    : 20,
    }

    m["Nuevo"] = 30

    fmt.Println(m)
    fmt.Println(m["Batman"])

    // When a key isn't defined, go return 0
    fmt.Println(m["AG"])

    // We can use "idioma coma okey"
    if v, ok := m["AG"]; ok {
        fmt.Printf("El valor desde if es: %v\n", v)
    }

    for i, v := range m {
        fmt.Printf("Llave: %v - Valor: %v\n", i, v)
    }

    if v, ok := m["AG"]; ok {
        fmt.Printf("Se borro la llave AG con valor: %v\n", v)
        delete(m, "AG")
    }

    fmt.Println(m)
    
}
