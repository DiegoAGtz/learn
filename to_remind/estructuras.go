/*
    Estructuras nos permiten crear tipos de datos compuestos
    Podemos embeber un struct en otro struct

    En go, la programación orientada a objetos no es tradicional. No existe la herencia como tal,
    para ello, usamos structs embebidos y el polimorfismo lo logramos mediante interfaces
*/
package main

import "fmt"

type persona struct {
    nombre string
    apellido string
    edad int
}

/*
    Cuando embebemos un struct, los tipos del struct embebido pasan a ser
    también del struct padre por lo que podemos hacer estudiante.nombre
    También funciona estudiante.persona.nombre, pero es mejor utilizar
    estudiante.nombre
    Los campos promovidos son los campos del struct embebido
*/

type estudiante struct {
    persona
    // si ponemos p1 persona
    // tenemos que usar estudiante.p1.nombre para acceder a sus campos
    nua int
    // x, y int
}

func main() {
    p1 := persona {
        nombre: "Diego",
        apellido: "Gutiérrez",
        edad: 20,
    }
    p2 := persona {
        nombre: "Armando",
        apellido: "Ayala",
        edad: 21,
    }

    e1 := estudiante {
        persona: persona {
            nombre: "Estudiante",
            apellido: "Apellido",
            edad: 33,
            /*
                Es posible usar
                    "Estudiante",
                    "Apellido",
                    33,
                Sin embargo es mejor usar el nombre de los campos, para mayor legilibidad
            */
        },
        nua: 147151,
    }

    fmt.Println(p1)
    fmt.Println(p2)
    fmt.Println(e1)

    // Acceso a los campos
    fmt.Println(p1.nombre)

    fmt.Println(e1.nombre)

    // Struct anonimo, es decir, sin identificador
    anonimo := struct {
        first string
        last string
        age int
    }{
        first: "Diego",
        last: "Gutierrez",
        age: 20,
    }
    fmt.Println(anonimo)
}
