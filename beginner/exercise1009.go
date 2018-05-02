package main

import (
	"fmt"
)

func main() {

	var name string
	var fixedSalary, sellingNumber float64

	fmt.Scan(&name, &fixedSalary, &sellingNumber)

	totalSalary := fixedSalary + 0.15*sellingNumber

	fmt.Printf("TOTAL = R$ %.2f\n", totalSalary)
}
