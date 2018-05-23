package main

import "fmt"

func main() {
	var k = 5

	for x := 1; x < 10; x += 2 {
		k = k + 2
		for y := k; y > k-3; y-- {
			fmt.Printf("I=%d J=%d\n", x, y)
		}
	}
}
