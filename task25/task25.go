package main

import (
	"fmt"
	"time"
)

//Реализовать собственную функцию sleep

func sleep(d time.Duration) {
	//time.After истечения продолжительности d, а затем отправляет текущее время по возвращаемому каналу
	<-time.After(d)
}

func main() {
	sleep(1 * time.Second)
	fmt.Println("a ...any")
}
