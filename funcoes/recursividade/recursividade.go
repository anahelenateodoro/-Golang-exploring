package main

import "fmt"

func fatorial(n uint) uint { //usar o uint n permite o uso de numeros negativos(-3,-2,-1)
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n-1)
	}
}

func main() {
	resultado := fatorial(5)
	fmt.Println(resultado)
}
