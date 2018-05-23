package main

import "fmt"

func main() {
	var x, y float64
	fmt.Scan(&x)
	for x < 0 || x > 10 {
		fmt.Print("nota invalida\n")
		fmt.Scan(&x)
	}
	if x >= 0 || x <= 10 {
		fmt.Scan(&y)
		for y < 0 || y > 10 {
			fmt.Print("nota invalida\n")
			fmt.Scan(&y)
		}
		if y >= 10 || y <= 10 {
			M := (x + y) / 2
			fmt.Printf("media = %.2f\n", M)
		}
	}
}
