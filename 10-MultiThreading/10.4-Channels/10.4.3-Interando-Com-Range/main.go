package main

import "fmt"

// Thread 1
func main() {
	ch := make(chan int)

	go publich(ch)
	// reader(ch)
	for x := range ch {
		fmt.Printf("Received %d\n", x)
	}
}

func reader(ch chan int) {
	for x := range ch {
		fmt.Printf("Received %d\n", x)
	}
}

func publich(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch) // Fechando o canal
}
