package main

import (
    "fmt"
    "os"
)

func main() {
    var (
        s1 string
        s2 string
        s3 string
    )

    n, err := fmt.Fscanln(os.Stdin, &s1, &s2, &s3)

    if err != nil {
        panic(err)
    } else {
        fmt.Printf("Se leyeron %v mensajes, su valor es: \t%v\n",n ,s1)
    }
}
