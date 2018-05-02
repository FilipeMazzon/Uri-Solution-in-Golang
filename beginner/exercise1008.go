package main

import (
	"fmt"
)

func main() {

	var number, hoursWorks int
	var valorHora float64

	//enter of data
	fmt.Scan(&number, &hoursWorks, &valorHora)
	//doing stuff
	salary := float64(hoursWorks) * valorHora

	fmt.Printf("NUMBER = %d\n", number)
	fmt.Printf("SALARY = U$ %.2f\n", salary)
}
