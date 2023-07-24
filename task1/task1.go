package main

import "fmt"

//Дана структура Human (с произвольным набором полей и методов).
//Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).

//структура Human
type Human struct {
	Height float32
	Sex    bool
}

//методы структуры
func (h *Human) Talk() {
	fmt.Println("bla bla bla")
}
func (h *Human) PrintHeigth() {
	s := fmt.Sprintf("%.2f", h.Height)
	fmt.Println(string(s))
}

//Встраивание полей структуры Human в Action
type Action struct {
	Human
}

func main() {
	//объявляем экземпляр Action, внутри него экземпляр Human
	act := Action{Human{Height: 180.2, Sex: false}}

	//методы сработают и так, и если вызвать через act.Human.Talk()
	act.PrintHeigth()
	act.Talk()
}
