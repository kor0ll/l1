package main

import "fmt"

//Разработать программу, которая переворачивает подаваемую на ход строку
//(например: «главрыба — абырвалг»). Символы могут быть unicode.

func main() {
	str := "главрыба"
	var result string

	//for range итерируется по rune, а не по byte
	for _, v := range str {
		//добавляем символ всегда в начало строки
		result = string(v) + result
	}
	fmt.Println(result)
}
