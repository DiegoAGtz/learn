/*
    Una función que se llama a si misma
*/

package main

import "fmt"

func main() {
    n1 := factorial(4)
    fmt.Println("Factorial:", n1)

    n2 := cicloFac(5)
    fmt.Println("Factorial:", n2)
}

func factorial(n int) int {
    if(n < 2) {
        return 1
    }
    return n * factorial(n - 1)
}

func cicloFac(n int) int {
    total := 1
    for ; n > 0; n-- {
        total *= n
    }
    return total
}
