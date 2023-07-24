package main

import (
	"fmt"
	"math"
)

//Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point
//с инкапсулированными параметрами x,y и конструктором.

type Point struct {
	x int
	y int
}

// методы для получения значений инкапсулированных параметров
func (p *Point) GetXPosition() int {
	return p.x
}
func (p *Point) GetYPosition() int {
	return p.y
}

// метод для расчета расстояния
func (p *Point) FindDistance(b *Point) float64 {
	//расстояние между точками рассчитывается по формуле AB = √(xb — xa)**2 + (yb — ya)**2
	return math.Sqrt(math.Pow(float64(b.GetXPosition())-float64(p.y), 2) + math.Pow(float64(b.GetYPosition())-float64(p.y), 2))
}

// конструктор для Point
func NewPoint(x int, y int) *Point {
	return &Point{x, y}
}

func main() {
	point1 := NewPoint(1, 1)
	point2 := NewPoint(3, 3)
	fmt.Println(point1.FindDistance(point2))
}
