package main

import (
	"fmt"
	"introducao-teste/enderecos"
)

func main() {
	typeAddress := enderecos.TypeAddress("Rua dos bobos")

	fmt.Println(typeAddress)
}
