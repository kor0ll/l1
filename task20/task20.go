package main

import (
	"fmt"
	"strings"
)

//Разработать программу, которая переворачивает слова в строке.
//Пример: «snow dog sun — sun dog snow».

func main() {
	str := "snow dog sun"
	//получаем слайс из слов
	s := strings.Split(str, " ")

	//записываем в другой слайс в обратном порядке
	reverseS := []string{}
	for i := len(s) - 1; i >= 0; i-- {
		reverseS = append(reverseS, s[i])
	}
	//склеиваем в строку
	reverseStr := strings.Join(reverseS, " ")
	fmt.Println(reverseStr)

}
