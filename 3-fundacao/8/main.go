// Funções, retorno de erro e múltiplos retornos

package main

import "fmt"

func main() {
	fmt.Printf("O resultado da soma é %d\n", soma(10, 20))

	soma, err := somaPositivos(12, 19)
	if err != nil {
		fmt.Printf("Erro: %s\n", err)
		return
	}
	fmt.Printf("O resultado da soma é %d\n", soma)

	soma, err = somaPositivos(-12, 19)
	if err != nil {
		fmt.Printf("Erro: %s\n", err)
		return
	}
}

func soma(a, b int) int {
	return a + b
}

func somaPositivos(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, fmt.Errorf("Somente números positivos são permitidos")
	}

	return a + b, nil
}
