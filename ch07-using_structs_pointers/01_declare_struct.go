package main

import "fmt"

type Movie struct {
	Name   string
	Rating float32
}

func main() {
	m := Movie{
		Name:   "R for Rajkumar",
		Rating: 10}
	fmt.Println(m.Name, m.Rating)

	var m2 Movie

	fmt.Println("m2 name is " + m2.Name)

	fmt.Printf("%+v\n", m2)

	m2.Name = "Metropolis"
	m2.Rating = 0.99

	fmt.Printf("%+v\n", m2)

	m3 := new(Movie)
	m3.Name = "Metropolis"
	m3.Rating = 0.99
	fmt.Printf("%+v\n", m3)
}
