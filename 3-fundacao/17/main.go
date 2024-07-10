// Ponteiro e Struct

package main

import "fmt"

type Cliente struct {
	Nome string
}

func (c *Cliente) andou() {
	c.Nome = "Jean Carlos Danese"
	fmt.Printf("%v andou.\n", c.Nome)
}

func main() {
	jean := Cliente{
		Nome: "Jean",
	}

	jean.andou()

	fmt.Printf("O nome do cliente é %v.\n", jean.Nome)
}
