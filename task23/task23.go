package main

import "fmt"

//Удалить i-ый элемент из слайса.

func main() {
	index := 2
	sl := []int{1, 2, 3, 4, 5}

	//сдвигаем значения влево на 1 индекс
	copy(sl[index:], sl[index+1:])

	//усекаем слайс
	sl = sl[:len(sl)-1]

	fmt.Println(sl)
}
