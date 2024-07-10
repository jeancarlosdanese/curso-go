// Funções variádicas

package main

func main() {
	println(soma(1, 2, 3, 4, 5))
	println(soma(1, 2, 3))
	println(soma(1, 2))
	println(soma(1))
	println(soma())
}

func soma(numeros ...int) int {
	soma := 0
	for _, numero := range numeros {
		soma += numero
	}
	return soma
}
