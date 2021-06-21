package main

import "fmt"

func main() {
    a := 42
    b := a << 1
    fmt.Printf("Decimal: %v\tBinario: %b\t\tHexadecimal: %X\n", a, a, a)
    fmt.Printf("Decimal: %v\tBinario: %b\tHexadecimal: %X\n", b, b, b)
}
