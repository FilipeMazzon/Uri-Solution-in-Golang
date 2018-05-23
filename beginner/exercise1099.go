package main

import "fmt"

func main() {

	var x, y, lol, a, b int
	var c = 0
	fmt.Scan(&lol)
	for d := 0; d < lol; d++ {
		fmt.Scan(&x, &y)
		if x > y {
			a = x
			b = y
		} else if x < y {
			a = y
			b = x
		}
		b++
		for ; b < a; b++ {
			if b%2 != 0 {
				c += b
			}
		}
		fmt.Printf("%d\n", c)
		c = 0
	}
}
