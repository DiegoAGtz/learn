package main

import "fmt"

func main() {
    n1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
    n2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    
    fmt.Println(foo(n1...))
    fmt.Println(bar(n2))
}

func foo(x ...int) int {
    suma := 0
    for _, v := range x {
        suma += v
    }
    return suma
}

func bar(x[] int) int {
    suma := 0
    for _, v := range x {
        suma += v
    }
    return suma
}
