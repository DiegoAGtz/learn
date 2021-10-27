package main

import (
	"encoding/json"
	"fmt"
)

type usuario struct {
    Nombre string
    Edad int
}

func main() {
    u1 := usuario {
        Nombre: "Diego",
        Edad: 20,
    }
    u2 := usuario {
        Nombre: "Ana",
        Edad: 19,
    }
    u3 := usuario {
        Nombre: "Ani",
        Edad: 20,
    }
    usuarios := []usuario{u1, u2, u3}
    fmt.Println(usuarios)

    bs, err := json.Marshal(usuarios)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(bs)
    fmt.Println(string(bs))
}
