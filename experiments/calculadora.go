/*
    Este programa realiza la suma de dos o más números introducids
    desde la linea de comandos.
    El programa se encarga de realizar algunas comprobaciones en los
    datos insertados de modo que únicamente acepta numeros enteros y
    da igual si, por ejemplo, se deja un espacio entre operaciones o
    se realiza sin espacio alguno.
*/
package main

import (
    "fmt"
    "os"
    "bufio"
    "strconv"
    "strings"
)

func calculadora() {

    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("Ingrese una operación, por ejemplo (2 + 3): ")

    scanner.Scan()

    s1 := scanner.Text()
    var operador string

    valores := strings.Split(s1, "+")
    operador = "+"
    if len(valores) < 0 {
        valores = strings.Split(s1, "-")
        operador = "-"
    }
    if len(valores) < 0 {
        valores = strings.Split(s1, "*")
        operador = "*"
    }
    if len(valores) < 0 {
        valores = strings.Split(s1, "/")
        operador = "/"
    }

    res := obtenerValor(valores[0])

    for i := 1; i < len(valores); i++ {
        res = operacion(res, obtenerValor(valores[i]), operador) 
    }

    fmt.Printf("El resultado de la operación es: %v\n", res)
}

func obtenerValor(cadena string) float64 {
    cadena = strings.Trim(cadena, " ")
    val, err := strconv.ParseFloat(cadena, 64)
    if err != nil {
        panic(err)
    }
    return val
}

func operacion(val1 float64, val2 float64, operando string) float64 {
    var res float64
    switch operando {
    case "+":
        res = val1 + val2
    case "-":
        res = val1 - val2
    case "*":
        res = val1 * val2
    case "/":
        res = val1 / val2
    }
    return res
}
