package main

import "fmt"

func main() {
    f := func() {
        fmt.Println("Hola desde la función anónima")
        for i:=0; i<10; i++ {
            fmt.Println(i)
        }
    }

    f()
    fmt.Printf("%T\n", f)
}
