package dirForStudy

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

/*
Заставить упасть родительскую горутину с паникой в дочерней, не делать
рекавери в дочерней и в родительской. Потом сделать правильный отлов
паники дочерней горутины.
*/

func RecoveryPanic() {
	errChan := make(chan interface{})

	go func() {
		go func() {
			defer func() {
				if r := recover(); r != nil {
					errChan <- r
				}
			}()
			breaking()
		}()
	}()

	select {
	case err := <-errChan:
		panic(err)
	case <-time.After(1 * time.Second):
		fmt.Println("main: всё прошло нормально")
	}
}

func breaking() {
	panic("что-то пошло не так")
}

/*
Разобраться как обработать сообщение, которое блокирует буферизированный
канал, не заблокировав его.
*/

func HandleBlock() {
	c := make(chan int, 3)

	for i := 0; i < 5; i++ {
		select {
		case c <- i:
			fmt.Printf("в канал записано %d\n", i)
		default:
			fmt.Printf("канал переполнен\n")
		}
	}
}

/*
Нужно ли явно закрывать канал? Будут ли утечки памяти, если его не закрыть?

Нет, закрывать канал нет необходимости в любой ситуации, и утечек он не вызывает. Случаи при которых нужно
закрыть канал это:
	1) показать, что больше не будет данных на запись;
	2) использовать := range для чтения из канала, иначе он не завершится, ожидая запись.
*/

/*
Написать реализацию принудительного переключения горутин во время
работы после длительной работы. runtime.Gosched()
*/

func SwitchScheduler() {
	// для имитации конкуренции за лог. ядро cpu
	runtime.GOMAXPROCS(1)
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count := 0
			for {
				count++
				fmt.Printf("выполняется горутина номер %d, итерация %d\n", i, count)
				if count >= 100_000 {
					break
				} else if count%1000 == 0 {
					runtime.Gosched()
				}
			}
		}()
	}

	wg.Wait()
}
