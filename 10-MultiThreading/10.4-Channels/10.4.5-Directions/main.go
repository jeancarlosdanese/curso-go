package main

import "fmt"

func receive(name string, ch chan<- string) {
	ch <- name
}

func read(ch <-chan string) {
	fmt.Println(<-ch)
}

// Thread 1
func main() {
	ch := make(chan string)

	// Thread 2
	go receive("Hello", ch)

	// Thread 3
	read(ch)
}
