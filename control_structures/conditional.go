package main

import "fmt"

func main() {
    // if condition { } 
    i := 10
    if i < 15 {
        fmt.Println("La condición es verdadera")
    }

    // We can reduce the scope of variables
    if x := 3; x <= 3 {
        fmt.Println(x)
    }
    // fmt.Println(x) x isn't defined

    // switch value <- If the value isn't defined, true is used

    switch {
    case 4 == 2, 3 == 1, 5 == 5:
        fmt.Println("NO, NO, SI imprime")
        fallthrough // Execute the next case even if it's false
    case 9 == 0:
        fmt.Println("No imprime")
    case 12 == 3:
        fmt.Println("No imprime")
    default:
        fmt.Println("Imprimiendo el default")
    }
    
    switch "Manzana" {
    case "Pera", "Aguacate", "Limon":
        fmt.Println("Frutas de color verde.")
    case "Manzana", "Fresa":
        fmt.Println("Frutas de color rojo.")
    default:
        fmt.Println("Frutas de otro color.")
    }
}
