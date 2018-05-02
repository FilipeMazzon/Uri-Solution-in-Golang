package main

import (
	"fmt"
)

func main()  {
	var a,b float64
	fmt.Scan(&a,&b)
	a*=3.5
	b*=7.5
	MEDIA := (a + b) / 11
	fmt.Printf("MEDIA = %.5f\n", MEDIA)
}