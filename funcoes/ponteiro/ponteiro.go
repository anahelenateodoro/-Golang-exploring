package main

import "fmt"

func inc1(n int) {
	n++ // n = n + 1
}

// revisão: um ponteiro é representado por um *
func inc2(n *int) {
	// revisap: * é usado para desferenciar, ou seja
	// ter acesso ao valor no qual o ponteiro aponta
	*n++
}

func main() {
	n := 1

	inc1(n) //por valor
	fmt.Println(n)

	// revisao: & usado para obter o endereço da variavel
	inc2(&n) // por referência
	fmt.Println(n)
}
