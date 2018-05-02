package main

import (
	"fmt"
	"math"
)

func highAmong(a, b, c float64) float64 {
	a = float64(a)
	largeAB := (a + b + math.Abs(a-b)) / 2
	largeABC := (largeAB + c + math.Abs(largeAB-c)) / 2
	return largeABC
}

func main() {
	var a, b, c int
	fmt.Scan(&a, &b, &c)
	a1 := float64(a)
	b2 := float64(b)
	c3 := float64(c)
	higherAmongABC := highAmong(a1, b2, c3)
	fmt.Printf("%.f eh o maior\n", higherAmongABC)
}
