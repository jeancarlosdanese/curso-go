// Continuação de Ponteiro e Struct

package main

import "fmt"

type Conta struct {
	Nome  string
	Saldo int
}

func NewConta(nome string, saldo int) *Conta {
	return &Conta{
		Nome:  nome,
		Saldo: saldo,
	}
}

// simularEntrada não altera o saldo real (não usa ponteiro)
func (c Conta) simularEntrada(valor int) {
	c.Saldo += valor
	fmt.Printf("O saldo simulado de %v é %v.\n", c.Nome, c.Saldo)
}

// simularSaida não altera o saldo real (não usa ponteiro)
func (c Conta) simularSaida(valor int) {
	c.Saldo -= valor
	fmt.Printf("O saldo simulado de %v é %v.\n", c.Nome, c.Saldo)
}

// efetivarEntrada altera o saldo real (usa ponteiro)
func (c *Conta) efetivarEntrada(valor int) {
	c.Saldo += valor
	fmt.Printf("O saldo atualizado de %v é %v.\n", c.Nome, c.Saldo)
}

// efetivarSaida altera o saldo real (usa ponteiro)
func (c *Conta) efetivarSaida(valor int) {
	c.Saldo -= valor
	fmt.Printf("O saldo atualizado de %v é %v.\n", c.Nome, c.Saldo)
}

func main() {
	// NewConta é um construtor que retorna um ponteiro
	cantaJean := NewConta("Conta Jean", 1000)
	fmt.Printf("O saldo inicial de %v é %v.\n\n", cantaJean.Nome, cantaJean.Saldo)

	// simularEntrada não altera o saldo real
	cantaJean.simularEntrada(100)
	fmt.Printf("O saldo real de %v é %v.\n\n", cantaJean.Nome, cantaJean.Saldo)

	// simularSaida não altera o saldo real
	cantaJean.simularSaida(100)
	fmt.Printf("O saldo real de %v é %v.\n\n", cantaJean.Nome, cantaJean.Saldo)

	// efetivarEntrada altera o saldo real
	cantaJean.efetivarEntrada(200)
	fmt.Printf("O saldo real de %v é %v.\n\n", cantaJean.Nome, cantaJean.Saldo)

	// efetivarSaida altera o saldo real
	cantaJean.efetivarSaida(500)
	fmt.Printf("O saldo real de %v é %v.\n\n", cantaJean.Nome, cantaJean.Saldo)
}
