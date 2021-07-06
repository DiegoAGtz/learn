package main

import (
    "fmt"
    "strconv"
    "os"
    "bufio"
)

func main() {
    fmt.Println("Bienvenido al menú de ejecución de programas")
    fmt.Println("--------------------------------------------")

    for true {
        printMenu()
        selection, ok := getSelection();
        if ok == true {
            fmt.Println("Ingrese sólo numeros")
        } else {
            if selection == 0 {
                break
            }
            if selection < 0 || selection > 7 {
                fmt.Println("Ingrese una opción válida")
                continue
            }
            fmt.Println()
            fmt.Println()
            fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++")
            fmt.Println()
            fmt.Println()
            switch selection {
            case 1:
                calculadora()
            case 2:
                contadorPalabras()
            case 3:
                slides()
            case 4:
                fibonacci()
            case 5:
                fibonacciClosure()
            case 6:
                scan()
            case 7:
                scan2()
            }
            fmt.Println()
            fmt.Println()
            fmt.Println("++++++++++++++++++++++++++++++++++++++++++++++++++++")
            fmt.Println()
            fmt.Println()
        }
    }
    fmt.Println("Ejecución terminada")
    fmt.Println("--------------------------------------------")
}

func printMenu() {
    fmt.Println("\n", "Seleccione una de las siguientes opciones")
    fmt.Println("1.- Calculadora")
    fmt.Println("2.- Contador de Palabras")
    fmt.Println("3.- Generador de imágenes")
    fmt.Println("4.- Fibonacci")
    fmt.Println("5.- Fibonacci con closures")
    fmt.Println("6.- Scan")
    fmt.Println("7.- Scan 2")
    fmt.Println("0.- Salir",)
}

func getSelection() (int, bool) {
    fmt.Print("\n", "Seleccion: ")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()

    res, err := strconv.Atoi(scanner.Text())
    ok := false
    if err != nil {
        ok = true
    }
    return res, ok
}
