package main

import "fmt"

func main() {

	var averageSpeed, travelTime float64

	fmt.Scan(&travelTime, &averageSpeed)

	distance := averageSpeed * travelTime

	liters := distance / 12

	fmt.Printf("%.3f\n", liters)
}
