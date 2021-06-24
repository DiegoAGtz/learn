package main

import "fmt"

func main() {
    // Only exist for loop
    
    // for init; condition; post { } <- like a C for
    for i := 0; i <= 100; i++ {
        fmt.Println(i)
    }

    // for condition { } <- like a C while 
    j := 0
    for j < 10 {
        fmt.Println(j)
        j++
    }

    // for { } <- Infinite loop, like a C for(;;)
    k := 0
    for {
        k++
        if k > 10 {
            break
        }
        fmt.Println(k)
    }

    // Usign an array, slice or map
    m := map[string]string {
        "Manzana": "Roja",
        "Pera": "Verde",
        "Aguacate": "Verde",
        "Naranja": "Naranja",
        "Sandia": "Verde/Roja",
    }

    for i, v := range m {
        fmt.Printf("Clave: %v - Valor: %v\n", i, v)
    }
}
