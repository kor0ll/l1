package main

import (
	"fmt"
	"strconv"
)

//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func setBit(num int64, pos int, valueChange bool) int64 {
	if valueChange {
		//используем логическое ИЛИ здесь
		return num | (1 << pos)
	} else {
		//и логическое И НЕ здесь
		return num &^ (1 << pos)
	}
}

func main() {
	var num int64 = 210
	fmt.Println(strconv.FormatInt(num, 2))

	fmt.Println(strconv.FormatInt(setBit(num, 1, false), 2))

	fmt.Println(12 &^ 10)

}
