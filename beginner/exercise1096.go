package main

import "fmt"

func main() {

	for x := 1; x < 10; x += 2 {
		for y := 7; y > 4; y-- {
			fmt.Printf("I=%d J=%d\n", x, y)
		}
	}
}
