package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func fibonacci() {
	fmt.Println("Programa que calcula los primeros n números de la serie de fibonacci")
	fmt.Print("Ingrese la cantidad de elementos a calcular: ")

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	nElementos, err := strconv.Atoi(scan.Text())

	if err != nil || nElementos < 1 {
		fmt.Println("Cantidad erronea")
	} else {
		printSlice(calcFibonacci(nElementos))
	}
}

func calcFibonacci(nElementos int) []int {
	fibonacci := []int{0}

	if nElementos > 1 {
		fibonacci = append(fibonacci, 1)
	}

	for i := 2; i < nElementos; i++ {
		fibonacci = append(fibonacci, fibonacci[i-2]+fibonacci[i-1])
	}

	return fibonacci
}

func printSlice(slice []int) {
	for i, val := range slice {
		fmt.Printf("%v ", val)
		if i%10 == 0 && i > 0 {
			fmt.Println()
		}
	}
}
