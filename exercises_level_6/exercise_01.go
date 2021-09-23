package main

import "fmt"

func main() {
    fmt.Println("Función foo():", foo())
    s, n := bar()
    fmt.Println("Función bar():", s, n)
}

func foo() int {
    return 42
}

func bar() (string ,int) {
    s1 := "Hola mundo desde bar."
    n1 := 3435
    return s1, n1
}
