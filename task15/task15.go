package main

import (
	"fmt"
	"strings"
)

//К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?
//Приведите корректный пример реализации.

//var justString string
//func someFunc() {
//  v := createHugeString(1 << 10)
//  justString = v[:100]
//}

//func main() {
//  someFunc()
//}

// В данном примере использование глобальной переменной нецелесообразно, т.к к ней можно обращаться
// из любой функции, ее изменение может быть неожиданным, что приведет к ошибкам
func createHugeString(length int) string {
	var str strings.Builder
	for i := 0; i < length; i++ {
		str.WriteString("x")
	}

	return str.String()
}

func someFunc() string {
	return createHugeString(1 << 10)[:100]
}

func main() {
	justString := someFunc()
	fmt.Println(justString)
}
