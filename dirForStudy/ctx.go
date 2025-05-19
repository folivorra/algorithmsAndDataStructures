package dirForStudy

import (
	"context"
	"fmt"
	"time"
)

// пакет context позволяет сигнализировать прекращение работы и вызывать return

/*
контексты соблюдают наследовательную систему создания, то есть для создания первоначального
контекста существует функция Background() и затем от него мы уже наследуем остальные
*/

func ContextCancellation() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	// создаем таймаут который прокинем через обе горутины последовательно

	defer cancel() // для гарантированного освобождени ресурсов

	go worker(ctx)

	time.Sleep(7 * time.Second) // ждем завершения
	fmt.Println("main completed")
}

func worker(ctx context.Context) {
	go workerForWorker(ctx)

	for {
		if ctx.Err() != nil {
			fmt.Println("worker cancelled with error:", ctx.Err())
			return
		}
		fmt.Println("worker doing job...")
		time.Sleep(1 * time.Second)
	}
}

func workerForWorker(ctx context.Context) {
	for {
		if ctx.Err() != nil {
			fmt.Println("workerForWorker cancelled with error:", ctx.Err())
			return
		}
		fmt.Println("workerForWorker doing job...")
		time.Sleep(1 * time.Second)
	}
}
