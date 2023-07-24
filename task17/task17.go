package main

import "fmt"

//Реализовать бинарный поиск встроенными методами языка.

func binarySearch(sl []int, elem int) int {
	//проверки входных данных
	if len(sl) == 0 {
		return -1
	}
	if sl[len(sl)-1] < elem {
		return -1
	}
	//обозначаем центральную позицию
	i := len(sl) / 2

	if sl[i] == elem {
		return i
	}
	if elem > sl[i] {
		return i + binarySearch(sl[i:], elem)
	} else {
		return binarySearch(sl[:i], elem)
	}

}

func main() {
	sl := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	fmt.Println(binarySearch(sl, 9))
}
