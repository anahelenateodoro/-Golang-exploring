package main

import "fmt"

func main() {
	i := 1

	// Go não tem aritmetica de ponteiros

	var p *int = nil

	p = &i // pegando o endereço da variavel
	*p++   //desreferenciando (pegando o valor)
	i++

	fmt.Println(p, *p, i)

	var newVar int32 = 2

	fmt.Println("Print main:", newVar, &newVar)

	printReference(&newVar)
	printAsCopy(newVar)
	fmt.Println("Print main:", newVar, &newVar)

}

func printReference(value *int32) {
	*value = 8
	fmt.Println("Print reference", *value, value)
}
func printAsCopy(value int32) {
	value = 4
	fmt.Println("Print as copy:", value, &value)
}
