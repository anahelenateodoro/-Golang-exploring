package main

import "fmt"

func main() {
	arrayDeNumeros := [5]int{1, 2, 3, 4, 5}
	arrayDeNumeros[2] = 33
	fmt.Println(arrayDeNumeros)
}
