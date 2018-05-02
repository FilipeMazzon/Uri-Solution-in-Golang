package main

import (
	"fmt"
	"math"
)

func main() {

	var Raio float64
	fmt.Scan(&Raio)
	const PI = 3.14159
	VOLUME := 4.0 / 3 * PI * math.Pow(Raio, 3)
	fmt.Printf("VOLUME = %.3f\n", VOLUME)

}
