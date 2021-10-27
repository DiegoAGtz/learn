package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type persona struct {
    Nombre string
    Edad int
}

func main() {
    p1 := persona {
        Nombre: "Diego",
        Edad: 20,
    }
    p2 := persona {
        Nombre: "Ana",
        Edad: 20,
    }
    personas := []persona{p1, p2}
    fmt.Println(personas)
    err := json.NewEncoder(os.Stdout).Encode(personas)
    if err != nil {
        fmt.Println(err)
    }
}
