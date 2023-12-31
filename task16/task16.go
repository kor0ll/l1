package main

import "fmt"

//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func quickSort(sl []int, left int, right int) []int {

	//Создаем копии пришедших переменных, с которыми будем манипулировать дальше
	l := left
	r := right

	//Вычисляем центр, на которой будем опираться
	center := sl[(left+right)/2]
	//Цикл, начинающий саму сортировку
	for l < r {

		//Пропускаем все значения больше центра, которые находятся справа
		for sl[r] > center {
			r--
		}
		//Также пропускаем значения меньше центра, которые находятся слева
		for sl[l] < center {
			l++
		}

		//В случае, когда находим значение меньше центрального справа и больше центрального слева,
		//делаем проверку
		if l <= r {
			//И если условие true, то меняем ячейки друг с другом.
			sl[r], sl[l] = sl[l], sl[r]
			l++
			r--
		}
	}
	//в конце данного цикла for получается, что то значение, которое мы выбрали как center, будет стоять на своем месте
	//т.к все значения меньше него находятся слева, а все больше него - справа
	//далее запускаем функцию снова на левую и правую части массива
	if r > left {
		quickSort(sl, left, r)
	}

	if l < right {
		quickSort(sl, l, right)
	}

	return sl
}

func main() {
	lst := []int{5, 4, 1, 2, 0, 123, 1234, 32, 12, 2345, 99}
	lst = quickSort(lst, 0, len(lst)-1)
	fmt.Println(lst)
}
