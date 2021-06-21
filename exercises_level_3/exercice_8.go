package main

import "fmt"

func main() {
    switch {
    case 8 == 9, "Grey" == "Gris", 0 == 0:
        fmt.Println("Esta linea se debería imprimir")
    case 9 < 10:
        fmt.Println("Esta linea no se debería imprimir")
    }
}
