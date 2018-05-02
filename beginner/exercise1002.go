package main

import (
	"fmt"
	"math"
)

func main() {
	var Raio float64
	fmt.Scan(&Raio)
	const PI = 3.14159
	VOLUME := PI * math.Pow(Raio, 2)
	fmt.Printf("A=%.4f\n", VOLUME)
}
