// Pacotes e módulos parte 3

package main

import (
	"fmt"

	"curso-go/matematica"
)

func main() {
	soma := matematica.Soma(10, 20)
	fmt.Printf("A soma de 1 e 2 é %v.\n", soma)

	fmt.Println(matematica.A)

	carro := matematica.Carro{Marca: "Fiat"}
	fmt.Println(carro.Marca)

	fmt.Println(carro.Andar())
}
