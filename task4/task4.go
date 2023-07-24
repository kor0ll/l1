package main

import (
	"context"
	"fmt"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

//Реализовать постоянную запись данных в канал (главный поток).
//Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
//Необходима возможность выбора количества воркеров при старте.

//Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

const (
	CountWorkers = 3
)

// функция, вызывающая воркеров
func StartWorkers(count int, ch chan string, ctx context.Context, wg *sync.WaitGroup) {
	for i := 0; i < count; i++ {
		wg.Add(1)
		//каждый воркер принимает канал для чтения данных и его номер
		go func(ch chan string, num int, ctx context.Context, wg *sync.WaitGroup) {
			//воркер в бесконечном цикле ждет, чтобы прочитать данные из канала и вывести, либо завершить работу
			for {
				select {
				//по сигналу от контекста завершаем работу
				case <-ctx.Done():
					fmt.Println("Worker N", num+1, "завершил работу!")
					wg.Done()
					return
				case data := <-ch:
					fmt.Println("Worker №", num+1, "received:", data)
				}
			}
		}(ch, i, ctx, wg)
	}
}

func main() {
	//завершать работу воркеров буду через контекст c сигналом
	//с помощью signal.NotifyContext можно перехватить выход из программы с помощью ctrl+c
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT)
	defer stop()
	wg := sync.WaitGroup{}
	ch := make(chan string)

	//запускаем воркеров
	StartWorkers(CountWorkers, ch, ctx, &wg)

	//запускаем постоянную запись в канал
	wg.Add(1)
	go func(ch chan string, ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Finished!")
				wg.Done()
				return
			default:
				ch <- "Hello from main!"
				time.Sleep(300 * time.Millisecond)
			}
		}
	}(ch, ctx)

	wg.Wait()
}
