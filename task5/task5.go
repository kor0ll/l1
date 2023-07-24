package main

import (
	"context"
	"fmt"
	"time"
)

//Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
//По истечению N секунд программа должна завершаться.

func main() {
	//количество секунд
	var n time.Duration = 4

	//объявляем контекст для завершения работы через определенное количество времени
	ctx, cancel := context.WithTimeout(context.Background(), n*time.Second)
	defer cancel()

	ch := make(chan string)

	go func(c context.Context, ch chan string) {
		for {
			select {
			//как только контекст отработал, выходим
			case <-c.Done():
				return
			//иначе слушаем канал
			case data := <-ch:
				fmt.Println(data)
			}
		}
	}(ctx, ch)

	//цикл слушает контекст для завершения работы и пишет данные в канал
	for {
		select {
		case <-ctx.Done():
			return
		default:
			ch <- "hello"
		}
	}

}
