package main

import "fmt"

func main() {
	x:= 3
	y:=4
	z:=5

	x, y, z = z, z+y, z+y+x

	fmt.Printf("%d %d %d\n", x, y, z)
}
