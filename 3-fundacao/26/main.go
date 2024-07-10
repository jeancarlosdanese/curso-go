// Looping and Infinite Loops in Go

package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("O valor de i é %d\n", i)
	}

	numeros := []int{10, 20, 30, 40, 50}
	for i, v := range numeros {
		fmt.Printf("O valor de i é %d e o valor de v é %d\n", i, v)
	}

	for _, v := range numeros {
		fmt.Printf("O valor de v é %d\n", v)
	}

	count := 10
	for i := 0; i < count; i++ {
		fmt.Printf("O valor de i é %d\n", i)
	}

	for {
		fmt.Println("Loop infinito")
	}
}
