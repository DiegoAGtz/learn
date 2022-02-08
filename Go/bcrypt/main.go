package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
    clave := "1234"
    bs, err := bcrypt.GenerateFromPassword([]byte(clave), bcrypt.MinCost)
    if err != nil {
        fmt.Println(err)
    }
    fmt.Println(clave)
    fmt.Println(bs)
    fmt.Println(string(bs))

    clave_login := "1234"
    err = bcrypt.CompareHashAndPassword(bs, []byte(clave_login))
    if err != nil {
        fmt.Println("Contraseña invalida")
        return
    }
    fmt.Println("Login correcto")
}
