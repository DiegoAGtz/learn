/*
    Uso de structs anonimos
*/
package main

import "fmt"

func main() {
    mEstudiante := struct{
        nombre string
        edad int
        nua int
    }{
        nombre: "Diego Armando",
        edad: 20,
        nua: 147151,
    }
    fmt.Println(mEstudiante)
}
