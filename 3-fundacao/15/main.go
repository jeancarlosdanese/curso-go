// Ponteiros

// Ponteiros são endereços de memória

// & -> pega o endereço da variável
// * -> pega o valor da variável

package main

func main() {
	a := 1
	b := &a
	c := *b
	println(a)
	println(b)
	println(c)
}
