/*
   Determinan cuales métodos adjuntar a un tipo
   Los métodos con el receptor pueden recibir argumentos del tipo
   valor.
   Son usados de la mano con interfaces.

   Receptor    Valor
   No puntero  No puntero y Puntero
   Puntero     Puntero


   Receptor     Valor
   (t T)        T, *T
   (t *T)       *T
*/
package main

import (
	"fmt"
	"math"
)

type cuadrado struct {
    lado float64
}

type circulo struct {
    radio float64
}

type forma interface {
    area() float64
}

func main() {
    c1 := cuadrado {
        lado: 5,
    }

    c2 := circulo {
        radio: 1,
    }

    info(c1)
    info(&c1)

    // Utiliar un valor no puntero para un receptor puntero dará un
    // error
    // info(c2)
    info(&c2)
}

func (c cuadrado) area() float64 {
    return math.Pow(c.lado, 2)
}

func (c *circulo) area() float64 {
    return math.Pi * math.Pow(c.radio, 2)
}

func info(s forma) {
    fmt.Println("El área es:", s.area())
}
