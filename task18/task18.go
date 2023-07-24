package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
//По завершению программа должна выводить итоговое значение счетчика.

// структура counter
type Counter struct {
	num int32
}

// функция позволяющая инкрементировать значение num в конкурентной среде
func (c *Counter) Increment() {
	//для простых конкурентных операций вместо Mutex можно использовать atomic
	atomic.AddInt32(&c.num, 1)
}

func main() {
	counter := Counter{}
	wg := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			counter.Increment()
			wg.Done()
		}(&wg)
	}
	wg.Wait()
	fmt.Println(counter.num)
}
