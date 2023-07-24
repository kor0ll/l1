package main

import (
	"context"
	"fmt"
	"time"
)

//Реализовать все возможные способы остановки выполнения горутины.

// способов остановки горутины 3, один из них - завершение main функции
func main() {
	//второй - с помощью прослушивания канала
	exit := make(chan bool)

	//третий - через context
	ctx, cancel := context.WithCancel(context.Background())
	//горутина, прекращающая работу при помощи канала
	go func(ch chan bool) {
		for {
			select {
			case <-exit:
				fmt.Println("Горутина с каналом завершена!")
				return
			default:
				fmt.Println("Горутина с каналом работает!")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(exit)
	//горутина, прекращающая работу при помощи контекста
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Горутина с контекстом завершена!")
				return
			default:
				fmt.Println("Горутина с контекстом работает!")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(2 * time.Second)
	//завершаем горутину с каналом
	exit <- true
	time.Sleep(1 * time.Second)
	//завершаем горутину с контекстом
	cancel()
	//если убрать time.Sleep снизу, функция main завершится раньше, чем успеет завершиться горутина с контекстом
	time.Sleep(1 * time.Second)
}
