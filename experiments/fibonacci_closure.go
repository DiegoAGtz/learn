/*
	Los closures o "cierres" son implementaciones de las funciones anónimas.
	Debido a esta propiedad, es posible implementar las funciones como
	una única instancia.
	En este ejemplo, calculamos los primeros 10 elementos de la serie de fibonacci.
*/
package main

import "fmt"

// fibonacci es una función que retorna una función que retorna un entero
func calFibonacci() func() int {
	fib1, fib2 := 0, 1
	return func() int {
		fib := fib1
		fib1, fib2 = fib2, fib1+fib2
		return fib
	}
}

func fibonacciClosure() {
	// f es nuestra "instancia", va guardando el estado
	f := calFibonacci()
	for i := 0; i < 10; i++ {
		fmt.Printf("%v ", f())
	}
	fmt.Println()
}
