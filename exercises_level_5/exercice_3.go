/*
    Crea un nuevo tipo: vehículo
        El tipo subyacente es struct
        Los campos:
            puertas
            color
    Crea dos nuevos tipos: camión y sedán
        El tipo subyacente de cada uno de esos tipos es un struct
        Embebe el tipo vehículo en ambos camión y sedán
        Dale al camión el campo cuatroRuedas el cual será un booelano
        Dale al sedán el campo lujoso el cual será un booleano
    Con los structs vehículo, camión y sedán
        Usando un composite literal, crea un valor de tipo camión y asignale valor a los campos
        Usando un composite literal, crea un valor de tipo sedan y asígnale valor a los campos
    Imprime cada uno de los valores
    Imprime un solo valor de cada uno de esos valores
*/
package main

import "fmt"

type vehiculo struct {
    puertas int
    color string
}
type camion struct {
    vehiculo
    cuatroRuedas bool
}
type sedan struct {
    vehiculo
    lujoso bool
}

func main() {
    mCamion := camion {
        vehiculo: vehiculo {
            puertas: 2,
            color: "Rojo",
        },
        cuatroRuedas: true,
    }
    mSedan := sedan {
        vehiculo: vehiculo {
            puertas: 4,
            color: "Negro",
        },
        lujoso: true,
    }
    fmt.Println(mCamion)
    fmt.Println(mCamion.puertas)
    fmt.Println(mSedan)
    fmt.Println(mSedan.puertas)
}
