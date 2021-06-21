package main

import "fmt"

func main() {
    if x := 3; x < 5 {
        fmt.Println("El valor de x es inferior a 5")
    } else if x > 5 {
        fmt.Println("El valor de x es superior a 5")
    } else {
        fmt.Println("El valor de x es 5")
    }
}
