package main

import "fmt"

func main() {

	var a, b uint64
	//don't working
	//este codigo está dando runtime error so .... 
	for {
		fmt.Scan(&a, &b)

		fmt.Printf("%d\n", a^b)
	}

}
