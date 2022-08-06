package main

import "fmt"

func main() {
	boolTrue := 40 > 30
	boolFalse := 5 < 1
	comparation1 := boolFalse == boolFalse
	comparation2 := boolTrue == boolTrue
	fmt.Println(boolTrue, boolFalse, comparation1, comparation2)
}
