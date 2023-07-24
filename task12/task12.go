package main

import "fmt"

//Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func main() {
	sl := []string{"cat", "cat", "dog", "cat", "tree"}
	m := make(map[string]int)

	for _, v := range sl {
		//если в карте еще нет множества v - создаем его
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {
			//иначе инкрементируем количество
			m[v]++
		}
	}

	fmt.Println(m)
}
