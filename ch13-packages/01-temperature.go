package temperature

func CtoF(c float64) float64 {
	return (c * (9 / 5)) + 32
}

func FtoC(f float64) float64 {
	return (9 / 5) * (f - 32)
}
