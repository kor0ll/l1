package main

//Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
//Функция проверки должна быть регистронезависимой.
//Например:
//abcd — true
//abCdefAaf — false
//aabcd — false

import (
	"fmt"
	"strings"
)

func CheckUnique(str string) bool {
	//приводим строку к нижнему регистру
	lowerString := strings.ToLower(str)

	m := make(map[rune]bool)

	for _, char := range lowerString {
		//если в мапе нет этого символа, добавляем
		if _, ok := m[char]; !ok {
			m[char] = true
			continue
		}
		return false
	}
	return true
}

func main() {

	fmt.Println(CheckUnique("abcd"))
	fmt.Println(CheckUnique("abCdefAaf"))
	fmt.Println(CheckUnique("aabcd"))
}
