/*
   Un puntero es una dirección de memoria donde hay almacenado un valor
   de algún tipo.

   Todo en GO es pasado por valor, similar a c y c++

   Usar punteros si tenemos una variable con una gran cantidad de datos, 
   en lugar de estar pasando valores es más eficiente pasar la dirección
   de memoria
*/

package main

import (
	"fmt"
)

func main() {
    a := 42
    fmt.Println(a)
    fmt.Println(&a)
    fmt.Printf("%T, %T\n",a, &a)

    var b *int = &a
    fmt.Println(b)
    // Acceder al valor almacenado en memoria, desrreferencia
    fmt.Println(*b)

    *b = 43
    fmt.Println(a)
    fmt.Println(*b)

    x := 0
    fmt.Println("X antes:", x)
    fmt.Println("X dirección:", &x)
    foo(&x)
    fmt.Println("X después:", x)
}

func foo(y *int) {
    fmt.Println("Y antes:", *y)
    fmt.Println("Y dirección:", y)
    *y = 42
    fmt.Println("Y después:", *y)
}
