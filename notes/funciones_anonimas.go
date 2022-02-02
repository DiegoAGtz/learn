package main

import "fmt"

func main() {
    foo()

    func() {
        fmt.Println("La función anónima se ejecuto")
    }()

    func(x int) {
        fmt.Println("EL significado de la vida es: ", x)
    }(42)
}

func foo() {
    fmt.Println("foo se ejecuto")
}
