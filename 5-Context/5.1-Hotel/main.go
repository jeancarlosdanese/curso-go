package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	select {
	case <-ctx.Done():
		fmt.Println("Cancelando reserva de hotel")
		return
	case <-time.After(2 * time.Second):
		fmt.Println("Reserva de hotel realizada")
	}
}
