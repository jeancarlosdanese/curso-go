// Importação do pacote fmt e tipagem de variáveis.

package main

import "fmt"

const message = "Hello, World!"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Jean"
	e float64 = 12.4
	f ID      = 1
)

func main() {
	fmt.Printf("O tipo de b é %T\n", b)
	fmt.Printf("O tipo de c é %T\n", c)
	fmt.Printf("O tipo de d é %T\n", d)
	fmt.Printf("O tipo de e é %T\n", e)

	fmt.Printf("O tipo de f é %T\n", f)
}
