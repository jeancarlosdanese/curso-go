package main

import (
	"fmt"
	"time"
)

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(1 * time.Second)
	}
}

// Thread 1
func main() {
	data := make(chan int)

	// // Thread 2
	// go worker(1, data)
	// go worker(2, data)

	qtdWorkers := 1000000
	for i := 0; i < qtdWorkers; i++ {
		go worker(i, data)
	}

	for i := 0; i < 10000000; i++ {
		data <- i
	}
}
