package main

import (
	"fmt"
)

func main() {
	var month, day,totalDay int
	daysInMonth := [12]int{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 25}

	for {
		totalDay=0
		fmt.Scan(&month, &day)
		if month == 12 && day == 25 {
			fmt.Print("E natal!\n")
		} else if month == 12 && day == 24 {
			fmt.Print("E vespera de natal!\n")
		} else if month == 12 && day > 25 {
			fmt.Print("Ja passou!\n")
		} else {
			for i := month-1 ; i< 11; i++ {
				totalDay+= daysInMonth[i]
			}
			totalDay+= 25-day
			fmt.Printf("Faltam %d dias para o natal!\n",totalDay)
		}
	}
}
