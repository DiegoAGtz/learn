package main

import "fmt"

func main() {
    for i := 0; i <= 10000; i++ {
        fmt.Printf("%v ", i)
        if i % 10 == 0 {
            fmt.Println("")
        }
    }
}
