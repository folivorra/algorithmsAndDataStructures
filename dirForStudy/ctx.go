package dirForStudy

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"math/rand"
	"os"
	"os/signal"
	"sync"
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

func ContextWorker() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second) // задаем таймаут
	wg := &sync.WaitGroup{}
	defer cancel() // гарантированное освобождение ресурсов

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go ctxWorker(ctx, i, wg) // запуск 5 воркеров с таймаутом в 3 секунды
	}

	wg.Wait()
}

func ctxWorker(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()
	timeoutDuration := time.Duration(rand.Intn(5)) * time.Second
	timeout := time.NewTimer(timeoutDuration) // создаем таймер который закроет канал по истечению 0-4 секунд
	for {
		select {
		case <-timeout.C:
			fmt.Printf("timeout for worker %d ❌\n", id) // воркер не успел по таймеру
			return
		case <-ctx.Done():
			fmt.Printf("worker %d finished his job ✅\n", id) // воркер успел выполнить работу
			return
		default:
			fmt.Printf("worker %d doing job...\n", id) // демонсстрация работы
			time.Sleep(1 * time.Second)
		}
	}
}

func GracefulShutdown() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt) // создание контекста который отлавливает прерывания на уровне ОС
	defer cancel()

	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go workerSignal(ctx, i, wg)
	}

	fmt.Println("program started, ctrl+c for finish")
	wg.Wait()
	fmt.Println("workers finished")
}

func workerSignal(ctx context.Context, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("worker %d stoped\n", id)
			return
			// ловим прерывание - прекращаем выполнение горутины и возвращаемся из функции
		default:
			fmt.Printf("worker %d doing job...\n", id)
			time.Sleep(1 * time.Second)
		}
	}
}

func ErrorGroup() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	group, ctx := errgroup.WithContext(context.Background())
	// функция создает группу и контекст для этой группы, который отменится если кто то из группы завершится с ошибкой

	for i := 0; i < 5; i++ {
		i := i
		// защита от замыкания
		group.Go(func() error {
			select {
			case <-time.After(time.Second * time.Duration(rand.Intn(5)+1)):
				if i == rand.Intn(5) {
					return fmt.Errorf("worker %d finished with error", i)
				}
				// имитируем случайную ошибку и возвращаем ее -> отменяем контекст группы

				fmt.Printf("worker %d done his job\n", i)
				// если ошибки не было значит воркер выполнил свою работу
				return nil
			case <-ctx.Done():
				// когда контекст группы отменяется - этот канал закрывается и мы возвращаем ошибку завершения группы
				return ctx.Err()
			}
		})
	}

	// ждем выполнение всех горутин в группе и если ошибка ненулевая, то выводим ее, иначе все ок
	if err := group.Wait(); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("worker finished without error")
	}
}
