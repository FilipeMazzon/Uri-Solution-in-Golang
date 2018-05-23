package main

import "fmt"

func main() {
	var cont int
	var a, b float64
	fmt.Scan(&cont)
	for i := 0; i < cont; i++ {
		fmt.Scan(&a, &b)
		if b == 0 {
			fmt.Print("divisao impossivel\n")
		} else {
			fmt.Printf("%.1f\n", a/b)
		}
	}
}
