package main

import (
	"fmt"
	"sync"
)

//Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

func main() {
	sl := [5]int{2, 4, 6, 8, 10}
	//объявляем wg, чтобы дождаться выполнения всех горутин
	wg := sync.WaitGroup{}
	var result int

	for _, value := range sl {
		wg.Add(1)
		go func(wg *sync.WaitGroup, val int) {
			//находим квадрат val и суммируем с result
			double := val * val
			result += double
			wg.Done()
		}(&wg, value)
	}
	//дожидаемся выполнения всех горутин
	wg.Wait()
	fmt.Println(result)
}
