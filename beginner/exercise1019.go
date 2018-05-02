package main

import "fmt"

func main() {
	var hora int
	fmt.Scan(&hora)

	fmt.Printf("%d:", hora/3600)
	hora %= 3600
	fmt.Printf("%d:", hora/60)
	fmt.Printf("%d\n", hora%60)

}
