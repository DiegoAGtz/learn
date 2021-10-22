// Funciones anónimas

package main

import "fmt"

func main() {
    func() {
        fmt.Println("Hola desde la función anónima")
        for i:=0; i<10; i++ {
            fmt.Println(i)
        }
    }()
}
