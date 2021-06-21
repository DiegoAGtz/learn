package main

import "fmt"

func main() {
    s1 := `Este es un strign de tipo raw string
            como vemos interpreta, incluso, los saltos
            de linea, tabulaciones y espacios. 
            También es llamado, "String literal no interpretado",
            también acepta comillas`

    fmt.Println(s1)
}
