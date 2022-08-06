package main

import "fmt"

func obertValorAprovado(numero int) int {
	defer fmt.Println("Fim!")
	if numero > 5000 {
		fmt.Println("Valor alto...")
		return 5000
	} // n precisa dee else
	fmt.Println("Valor baixo...")
	return numero
}

func main() {
	fmt.Println(obertValorAprovado(6000))
	fmt.Println(obertValorAprovado(3000))
}
