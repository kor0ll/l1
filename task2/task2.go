package main

import (
	"fmt"
	"sync"
)

//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	//объявляем wg, чтобы дождаться выполнения всех горутин
	wg := sync.WaitGroup{}

	for _, value := range arr {
		wg.Add(1)
		go func(wg *sync.WaitGroup, val int) {
			double := val * val
			fmt.Println(double)
			wg.Done()
		}(&wg, value)
	}
	//дожидаемся выполнения всех горутин
	wg.Wait()
}
