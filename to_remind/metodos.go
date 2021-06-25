/*
    Los métodos son funciones que adjuntamos a algún tipo.
    Podemos verlos como los métodos de un objeto en una clase, de hecho, pienso
    que es el equivalente en Go
*/
package main

import "fmt"

type persona struct {
    nombre string
    apellido string
}

type agenteSecreto struct {
    persona
    lpm bool
}

func main() {
    as1 := agenteSecreto {
        persona: persona {
            nombre: "Diego",
            apellido: "Gutiérrez",
        },
        lpm: true,
    }
    as2 := agenteSecreto {
        persona: persona {
            nombre: "Armando",
            apellido: "Ayala",
        },
        lpm: false,
    }
    fmt.Println(as1)
    fmt.Println(as2)

    as1.presentarse()
    as2.presentarse()
}

// firma de la función
// func (r receptor) identificador(parámetros) (retorno(s)) { código }
func (a agenteSecreto) presentarse() {
    fmt.Println("Hola yo soy el agente secreto", a.nombre, a.apellido)
}

