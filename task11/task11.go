package main

import "fmt"

//Реализовать пересечение двух неупорядоченных множеств.

func main() {
	//объявляем 2 множества
	sl1 := []int{1, 5, 3, 4, 6}
	sl2 := []int{7, 2, 4, 1, 6, 11}

	//пересечения записываем в result
	result := []int{}

	m := make(map[int]struct{})

	//добавляем в map все значения из первого множества
	for _, v := range sl1 {
		m[v] = struct{}{}
	}
	//если значение из второго множества уже есть в map, добавляем его в result
	for _, v := range sl2 {
		if _, ok := m[v]; ok {
			result = append(result, v)
		}
	}

	fmt.Println(result)
}
