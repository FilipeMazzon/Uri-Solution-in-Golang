package main

import "fmt"

func main() {
	var x, qt int

	fmt.Scan(&x)

	if x >= 1 && x <= 5 {
		switch x {
		case 1:
			fmt.Scan(&qt)
			fmt.Printf("Total: R$ %d.00\n", qt*4)
		case 2:
			fmt.Scan(&qt)
			fmt.Printf("Total: R$ %.2f\n", float64(qt)*4.5)
		case 3:
			fmt.Scan(&qt)
			fmt.Printf("Total: R$ %d.00\n", qt*5)
		case 4:
			fmt.Scan(&qt)
			fmt.Printf("Total: R$ %d.00\n", qt*2)
		case 5:
			fmt.Scan(&qt)
			fmt.Printf("Total: R$ %.2f\n", float64(qt)*1.5)
		}
	}
}
