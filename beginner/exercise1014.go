package main

import "fmt"

func main() {

	var fuelExpense float64
	var travelledDistance int

	fmt.Scan(&travelledDistance, &fuelExpense)

	averageConsumption := float64(travelledDistance) / fuelExpense;
	fmt.Printf("%.3f km/l\n", averageConsumption)
}
