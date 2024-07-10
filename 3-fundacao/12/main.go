// Composição de structs

package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {
	jean := Cliente{
		Nome:  "Jean",
		Idade: 33,
		Ativo: true,
	}
	fmt.Printf("O cliente %v tem %d anos.\n", jean.Nome, jean.Idade)

	fmt.Printf("O cliente %v está ativo? %v\n", jean.Nome, jean.Ativo)

	jean.Endereco = Endereco{
		Logradouro: "Rua das Flores",
		Numero:     123,
		Cidade:     "São Paulo",
		Estado:     "SP",
	}

	fmt.Printf("O cliente %v mora na %v, número %d, %v, %v\n", jean.Nome, jean.Endereco.Logradouro, jean.Endereco.Numero, jean.Endereco.Cidade, jean.Endereco.Estado)
}
