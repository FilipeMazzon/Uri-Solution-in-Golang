package main

import (
	"fmt"
)

func main() {
	var x float64
	fmt.Scan(&x)
	if x < 0 {
		fmt.Print("Fora de intervalo\n")
	} else if x <= 25 {
		fmt.Print("Intervalo [0,25]\n")
	} else if x <= 50 {
		fmt.Print("Intervalo (25,50]\n")
	} else if x <= 75 {
		fmt.Print("Intervalo (50,75]\n")
	} else if x <= 100 {
		fmt.Print("Intervalo (75,100]\n")
	} else {
		fmt.Print("Fora de intervalo\n")
	}
}
