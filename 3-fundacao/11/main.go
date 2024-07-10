// Iniciando com Structs

package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	cliente := Cliente{
		Nome:  "Jean",
		Idade: 33,
		Ativo: true,
	}
	fmt.Printf("O cliente %v tem %d anos.\n", cliente.Nome, cliente.Idade)

	fmt.Printf("O cliente %v está ativo? %v\n", cliente.Nome, cliente.Ativo)
}
