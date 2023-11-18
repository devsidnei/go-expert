package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Criar um contexto de fundo (background)
	ctx := context.Background()

	// Criar um contexto com cancelamento após 3 segundos
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	// Chamar uma função que utiliza o contexto
	doSomething(ctx)
}

func doSomething(ctx context.Context) {
	// Iniciar uma goroutine que realiza uma tarefa demorada
	go func() {
		// Simulando uma tarefa que leva 5 segundos
		time.Sleep(5 * time.Second)

		// Verificar se o contexto foi cancelado
		select {
		case <-ctx.Done():
			// O contexto foi cancelado antes da conclusão da tarefa
			fmt.Println("Tarefa cancelada!")
		default:
			// A tarefa foi concluída com sucesso
			fmt.Println("Tarefa concluída com sucesso!")
		}
	}()

	// Agora, você pode fazer outras coisas enquanto a tarefa está em andamento
	fmt.Println("Aguardando o término da tarefa...")

	select {
	case <-ctx.Done():
		// O contexto foi cancelado antes da conclusão da tarefa
		fmt.Println("Tarefa cancelada durante a espera.")
	case <-time.After(6 * time.Second):
		// Apenas para demonstração, espere por um tempo
		fmt.Println("Continuando após um tempo de espera.")
	}
}
