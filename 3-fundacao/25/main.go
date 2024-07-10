// Pacotes e módulos parte 3
// Compilar e executar o código
// go tool dist list

package main

import (
	"fmt"

	"curso-go/matematica"

	"github.com/google/uuid"
)

func main() {
	soma := matematica.Soma(10, 20)
	fmt.Printf("A soma de 1 e 2 é %v.\n", soma)

	fmt.Println(matematica.A)

	carro := matematica.Carro{ID: uuid.New(), Marca: "Fiat"}
	fmt.Println(carro.ID, carro.Marca)

	fmt.Println(carro.Andar())
}
