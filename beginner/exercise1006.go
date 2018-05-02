package main

import (
	"fmt"
)

func main() {
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	a *= 2
	b *= 3
	c *= 5
	MEDIA := (a + b + c) / 10
	fmt.Printf("MEDIA = %.1f\n", MEDIA)
}
