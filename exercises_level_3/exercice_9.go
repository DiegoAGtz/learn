package main

import "fmt"

func main() {
    deporteFav := "soccer"
    switch deporteFav {
    case "soccer", "basketball", "boleiball":
        fmt.Println("Deportes con pelota.")
    case "atletismo", "caminata":
        fmt.Println("Deportes de condición")
    case "boxeo", "lucha libre":
        fmt.Println("Deportes de fuerza")
    }
}
