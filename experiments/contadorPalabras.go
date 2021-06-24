package main
import (
    "fmt"
    "strings"
    "os"
    "bufio"
)

func main() {
    // frase := "Hola que tal, como están todos. Tal es el caso de este texto que usamos para probar nuestro programa"
    // frase := "Saludos a todos los presentes. Como están los presentes en este código?¿, además, vamos a saludar con muchos saludos a los que nos acompañan aquí aquí aquí"
    fmt.Println("Este es un contador de palabras.")
    fmt.Println("Ingrese un texto y el programa le contará la cantidad de palabras distintas en dicho texto:")
    frase := getMessage()

    split_frase := strings.Split(frase, " ")
    m := map[string]int{}
    
    for _, v := range split_frase {
        v = trim(v)
        if _, ok := m[v]; !ok {
            m[v] = 1
            continue
        }
        m[v]++
    }

    for k, v := range m {
        fmt.Println(k, "->", v)
    }
}

func trim(s string) string {
    var res string
    res = strings.ToLower(s)
    res = strings.Replace(res, ",", "", -1)
    res = strings.Replace(res, ".", "", -1)
    res = strings.Replace(res, "¡", "", -1)
    res = strings.Replace(res, "!", "", -1)
    res = strings.Replace(res, "¿", "", -1)
    res = strings.Replace(res, "?", "", -1)
    return res
}

func getMessage() string {
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    return scanner.Text()
}
