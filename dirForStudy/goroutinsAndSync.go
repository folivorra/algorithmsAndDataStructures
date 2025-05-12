package dirForStudy

import (
	"fmt"
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
