package main

import "fmt"

func main() {
	var players = make(map[string]int)
	players["cook"] = 32
	players["baistow"] = 27
	players["stokes"] = 26

	fmt.Println(players)

	// delete an element from map

	delete(players, "cook")
	fmt.Println(players)
}
