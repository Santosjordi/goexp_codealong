package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	bookHotel(ctx)
}

func bookHotel(ctx context.Context) {
	// imagine que aqui temos uma chamada para um serviço de reserva de hotel
	// que pode demorar um pouco para responder
	// o contexto é passado para a chamada para que possamos cancelar a operação
	// caso o usuário cancele a operação
	select {
	case <-ctx.Done():
		fmt.Println("Operação cancelada. Timeout excedido.")
		return
	case <-time.After(3 * time.Second):
		fmt.Println("Hotel reservado!")
		return
	}
}
