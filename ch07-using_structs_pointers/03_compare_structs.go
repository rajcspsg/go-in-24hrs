package main

import "fmt"
import "reflect"

type Drink struct {
	Name string
	Ice  bool
}

func main() {
	a := Drink{
		Name: "Lemonade",
		Ice:  true,
	}

	b := Drink{
		Name: "Lemonade",
		Ice:  true,
	}

	fmt.Println(" a == b ", a == b)
	fmt.Println(reflect.TypeOf(a))
}
