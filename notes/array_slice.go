package main

import "fmt"

func main() {
    // An array is static, requires a lenght
    // Arrays with different size, are different types
    var array1 [6]int
    array2 := [5]int {0, 1, 2, 3, 4}

    // A slice is dynamic, contains a subjacent array that isn't dynamic
    // A slice have lenght, which is the number of elements, and have capacity,
    // that is the lenght of the subjacent array
    var slice1 []int
    slice2 := []int {0, 1, 2, 3, 4}

    // make function limit the size of the subjacent array and asign a size
    // to slice
    // When the subjacent array is overflow, it double its size
    // this makes the execution more efficient
    slice3 := make([]int, 10, 12)

    slice1 = append(slice1, 19, 20, 39)
    slice3 = append(slice3, slice1...)

    slice1 = append(slice2[:1], slice2[3:]...)

    fmt.Println(array1)
    fmt.Println(array2)
    fmt.Println(slice1)
    fmt.Println(slice2)
    fmt.Println(slice3)

    fmt.Println(len(slice3))
    fmt.Println(cap(slice3))

    // Print with for loop
    for i, val := range slice3 {
        fmt.Printf("Iteración: %v, Valor: %v\n", i, val)
    }
}
