package main

import (
    "fmt"
    "os"
    "bufio"
)

func main() {

    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("Ingrese un mensaje: ")

    scanner.Scan()

    s1 := scanner.Text()

    fmt.Printf("El mensaje escrito es: \t%v\n", s1)
}
