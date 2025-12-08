package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Crição de um contexto com timeout de 6 segundos, ou seja, se a operação demorar mais que isso, ela será cancelada automaticamente
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*6)
	// Ao finalizar a função main, chamamos cancel para liberar recursos associados ao contexto
	defer cancel()

	bookHotel(ctx)

}

// Simulação de uma função que tenta reservar um quarto de hotel, respeitando o contexto fornecido
func bookHotel(ctx context.Context) {
	select {
		// Se o contexto for cancelado (por timeout ou manualmente), essa case será executada
		case <-ctx.Done():
			fmt.Println("Tempo Excedido para bookar o quarto")

		// Simulação de uma operação que leva 5 segundos para completar
		case <-time.After(5 * time.Second):
			fmt.Println("Quarto bookado com sucesso")
	}
}

