package main

import (
	"fmt"
	"math"
)

type Sphere struct {
	Radius float64
}

func (s *Sphere) SurfaceArea() float64 {
	return float64(4) * math.Pi * (s.Radius * s.Radius)
}

func (s *Sphere) Volume() float64 {
	return float64(4/3) * math.Pi * (s.Radius * s.Radius * s.Radius)
}

func main() {
	s := Sphere{Radius: 7}
	fmt.Println(s.SurfaceArea())
	fmt.Println(s.Volume())
}