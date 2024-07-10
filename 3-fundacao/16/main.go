// Quando usar ponteiros

package main

func main() {
	a := 1
	b := 2

	println(soma(&a, &b))
	println(a)
}

func soma(a, b *int) int {
	*a = 10
	return *a + *b
}
