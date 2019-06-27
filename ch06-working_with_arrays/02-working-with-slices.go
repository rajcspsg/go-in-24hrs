package main

import "fmt"

func main() {
	var cheeses = make([]string)
	cheeses[0] = "Mariolles"
	cheeses[1] = "Époisses de Bourgogne"

	fmt.Println(cheeses[0])
	fmt.Println(cheeses[1])
	fmt.Println(cheeses)

	fmt.Println(append(cheeses, "Camembert"))
	fmt.Println(cheeses)

	cheeses = append(cheeses, "Camembert")
	fmt.Println(cheeses)

	cheeses = append(cheeses, "Reblochon", "Picodon")
	fmt.Println(cheeses)

	// removing elements from slices
	cheeses = append(cheeses[:2], cheeses[2+1:]...)
	fmt.Println(cheeses)

	// copying the slices
	var smellyCheeses = make([]string, 4)
	copy(smellyCheeses, cheeses)
	fmt.Println(smellyCheeses)
}
