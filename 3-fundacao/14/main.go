// Interfaces

package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface {
	Desativar()
}

func Desativar(p Pessoa) {
	p.Desativar()
}

type Empresa struct {
	Nome  string
	Ativo bool
}

func (e *Empresa) Desativar() {
	e.Ativo = false
	fmt.Printf("A empresa %v foi desativada.\n", e.Nome)
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func (c *Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("O cliente %v agora está inativo.\n", c.Nome)
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

	jean.Desativar()

	Desativar(&jean)

	empresa := Empresa{
		Nome:  "Cetesc",
		Ativo: true,
	}

	Desativar(&empresa)
}
