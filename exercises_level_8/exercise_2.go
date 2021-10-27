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
    datos := `[{"Nombre":"Diego","Edad":20},{"Nombre":"Ana","Edad":19},{"Nombre":"Ani","Edad":20}]`
    bs := []byte(datos)
    var usuarios []usuario
    err := json.Unmarshal(bs, &usuarios)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(datos)
    fmt.Println(usuarios)
}
