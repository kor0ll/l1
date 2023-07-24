package main

import (
	"fmt"
	"reflect"
)

//Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

func checkType(i interface{}) string {
	//также можно через switch i.(type), но с каналом не сработает, т.к нужно обязательно указывать тип канала
	return reflect.TypeOf(i).String()
}

func main() {
	num := 5
	boolean := true
	ch := make(chan int)
	str := ""

	fmt.Println(checkType(num))
	fmt.Println(checkType(boolean))
	fmt.Println(checkType(ch))
	fmt.Println(checkType(str))
}
