package main

// Thread 1
func main() {
	forever := make(chan bool) // Canal vazio

	// forever <- true // Não funciona, precisa de uma goroutine

	go func() {
		for i := 0; i < 10; i++ {
			println(i)
		}
		forever <- true // Canal cheio
	}()

	<-forever // Canal esvaziado
}
