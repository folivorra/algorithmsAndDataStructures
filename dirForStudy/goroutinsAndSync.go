package dirForStudy

import (
	"fmt"
	"time"
)

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
