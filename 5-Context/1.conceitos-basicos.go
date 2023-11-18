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
		fmt.Println("O Hotel cancelou a reserva")
		return
	case <-time.After(1 * time.Second):
		{
			fmt.Println("Reserva realizada com sucesso!")
		}
	}
}
