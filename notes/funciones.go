/*
    func (r receptor) identificador(parámetros) retorno(s) { código }
    Cuando tenemos un receptor, se llama método
    Si no lo tenemos se llama función
    En Go, todo es pasado por valor, nada es pasado por copia ni por referencia
    como en otros lenguajes. WYSIWYG (what you see is what you get)

    cuando definimos la función se llama parametro
    cuando la llamamos la llamamos argumento
*/
package main

import "fmt"

func main() {
    fmt.Println("Debería imprimr el foo")
    defer foo() // Diferir -> Posponer, se ejecuta cuando la función que la contiene llega al su retorno o al final
    bar("Diego")
    s := woo("Armando.")
    fmt.Println(s)

    x, y := saludar("Diego", "Gutiérrez")
    fmt.Println(x, "-", y)

    xi := []int {1, 2, 3, 4, 5, 6, 7, 8, 9}
    // suma := sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
    suma := sum("Hola", xi...) // Despliegue
    fmt.Println("El valor de suma es:", suma)

    fmt.Println("Apenas se imprime la función foo del comienzo")
}

func foo() {
    fmt.Println("Hola desde foo.")
}

func bar(s string) {
    fmt.Println("Hola,", s)
}

func woo(s string) string {
    return fmt.Sprint("Hola desde woo,", s)
}

func saludar(n string, a string) (string, bool) {
    p := fmt.Sprint(n, " ", a, ` dice "hola".`)
    q := true
    return p, q
}

// Cantidad variable de parametros
// x es un slice de enteros
// Debe ser siempre el último parametro de la lista
func sum(s string, x ...int) int {
    suma := 0
    fmt.Println(x)
    fmt.Printf("%T\n", x)
    fmt.Println(len(x))
    fmt.Println(cap(x))
    fmt.Println(s)

    for i, v := range x {
        suma += v
        fmt.Println("El valor en el índice", i, "suma", v, "al total, quedando", suma)
    }
    return suma
}
