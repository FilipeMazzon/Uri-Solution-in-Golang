package main

import (
	"fmt"
	"time"
)

//don't working in this moment
func main() {
	var casoDeTeste, aux, aux2 int
	var word string
	fmt.Scan(&casoDeTeste)
	for i := 0; i < casoDeTeste; i++ {
		now := time.Now()
		fmt.Scan(&word)
		after := time.Now()
		aux = after.Second() - now.Second()

		if aux < 0 {
			aux *= -1
		}
		aux2 = after.Nanosecond() - now.Nanosecond()
		if aux2 < 0 {
			aux--
			aux2 *= -1
		}
		aux2 /= 10000000
		if aux2 < 10 {
			aux2 *= 10
		}
		if aux2 == 0 {
			fmt.Printf("%d.%d0\n", aux, aux2)
		} else {
			fmt.Printf("%d.%d\n", aux, aux2)
		}

	}

}
