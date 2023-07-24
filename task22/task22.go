package main

import (
	"fmt"
	"math/big"
)

//Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

func main() {

	//для математических операций с большими числами используется пакет big
	a := big.NewInt(5459454)
	b := big.NewInt(23423443)
	var result = new(big.Int)

	result.Add(a, b) //сложение
	fmt.Println(result)

	result.Mul(a, b) //Умножение
	fmt.Println(result)

	result.Sub(b, a) //Вычитание
	fmt.Println(result)

	result.Div(b, a) //Деление
	fmt.Println(result)

}
