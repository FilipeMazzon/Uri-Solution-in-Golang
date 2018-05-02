package main

import (
	"fmt"
)

func main() {
	var A [3] int

	for i := 0; i < 3; i++ {
		fmt.Scan(&A[i])
	}
	defer fmt.Printf("%d\n%d\n%d\n", A[0], A[1], A[2])
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if A[j] > A[i] {
				A[j], A[i] = A[i], A[j]
			}
		}
	}
	fmt.Printf("%d\n%d\n%d\n\n", A[0], A[1], A[2])
}
