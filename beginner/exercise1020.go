package main

import "fmt"

func main() {
	var data int
	fmt.Scan(&data)
	fmt.Printf("%d ano(s)\n", data/365)
	data %= 365
	fmt.Printf("%d mes(es)\n", data/30)
	data %= 30
	fmt.Printf("%d dia(s)\n", data)
}
