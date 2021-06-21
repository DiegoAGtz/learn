package main

import "fmt"

func main() {
    for i := 10; i <= 100; i++ {
        if i % 4 == 0 {
            fmt.Println("")
        }
        fmt.Printf("%v\t", i % 4)
    }
    fmt.Println("")
}
