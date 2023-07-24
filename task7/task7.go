package main

import "sync"

//Реализовать конкурентную запись данных в map.

type MyMap struct {
	mu sync.Mutex
	m  map[int]int
}

func NewMap() *MyMap {
	return &MyMap{
		mu: sync.Mutex{},
		m:  make(map[int]int),
	}
}

func (mp *MyMap) Store(key int, val int) {
	mp.mu.Lock()
	mp.m[key] = val
	mp.mu.Unlock()
}

func main() {
	m := NewMap()
	for i := 0; i < 1000; i++ {
		go func(i int) {
			m.Store(i, i+1)
		}(i)
	}
}
