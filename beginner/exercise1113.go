package main

import "fmt"

func main() {
	var m, n int

	fmt.Scan(&m, &n)
	for m != n {
		if m > n {
			fmt.Printf("Decrescente\n")
		} else {
			fmt.Printf("Crescente\n")
		}
		fmt.Scan(&m, &n)
	}
}
