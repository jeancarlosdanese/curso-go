// Declarando, atribuindo e acessando valores de um array

package main

import "fmt"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Jean"
	e float64 = 12.4
	f ID      = 1
)

func main() {
	var meuArray [3]int
	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 32

	fmt.Println(len(meuArray), meuArray[0], meuArray[1], meuArray[2])
	fmt.Println(meuArray[len(meuArray)-1])

	for i, v := range meuArray {
		fmt.Printf("O valor de i é %d e o valor de v é %d\n", i, v)
	}
}
