package main

import "fmt"

type Triangle struct {
	base   float64
	height float64
}

func (t *Triangle) area() float64 {
	return 0.5 * t.height * t.base
}

func (t Triangle) changeBase(f float64) {
	t.base = f
	return
}

func (t *Triangle) changeBase2(f float64) {
	t.base = f
	return
}

func main() {
	t := Triangle{base: 3, height: 1}
	fmt.Println(t.area())
	t.changeBase(4)
	fmt.Println(t.area())

	t.changeBase2(4)
	fmt.Println(t.area())
}
