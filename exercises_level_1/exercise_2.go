package main

import "fmt"

var x int
var y string
var z bool

func main() {
    fmt.Println(x, y, z)
    fmt.Println("")
    fmt.Println(holaMundo("Mensaje pasado como parametro"))
}

func holaMundo(mensaje string) string {
    s1 := fmt.Sprintf("Hola mundo desde una funcion con el mensaje: %s", mensaje)
    return s1 
}
