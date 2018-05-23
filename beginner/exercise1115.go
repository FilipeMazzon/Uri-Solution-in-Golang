package main

import "fmt"

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	for a != 0 && b != 0 {
		if a > 0 && b > 0 {
			fmt.Print("primeiro\n")
		} else if a > 0 && b < 0 {
			fmt.Print("quarto\n")
		} else if a < 0 && b > 0 {
			fmt.Print("segundo\n")
		} else {
			fmt.Print("terceiro\n")
		}
		fmt.Scan(&a, &b)
	}
}
