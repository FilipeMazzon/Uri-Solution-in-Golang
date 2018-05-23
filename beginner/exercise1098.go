package main

import "fmt"

func main() {

	var n1, k float64
	var q = 10

	for x := 0; x <= 20; x += 2 {
		n1 = float64(x) / 10.0

		for y := q; y < q+30; y += 10 {
			k = n1 + float64(y)/10.0
			if n1 == 0 || n1 == 1 || n1 == 2 {
				if k == 1 || k == 2 || k == 3 || k == 4 || k == 5 {
					fmt.Printf("I=%.F J=%.F\n", n1, k)
				} else {
					fmt.Printf("I=%.F J=%.1F\n", n1, k)
				}
			} else {
				if k == 1 || k == 2 || k == 3 || k == 4 || k == 5 {
					fmt.Printf("I=%.1F J=%.F\n", n1, k)
				} else {
					fmt.Printf("I=%.1F J=%.1F\n", n1, k)
				}
			}
		}
	}
}
