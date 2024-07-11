package main

func main() {
	ch := make(chan string, 2) // Buffer de 2 posições
	ch <- "Hello"
	ch <- "World"

	println(<-ch)
	println(<-ch)
}
