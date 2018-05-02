package main

import "fmt"

func main() {

	var valuePay, valueRecieve, n, n100, n50, n20, n10, n5, n2 int
	fmt.Scan(&valuePay, &valueRecieve)
	for valuePay+valueRecieve != 0 {
		n = valueRecieve - valuePay
		n100 = n / 100
		n %= 100
		n50 = n / 50
		n %= 50
		n20 = n / 20
		n %= 20
		n10 = n / 10
		n %= 10
		n5 = n / 5
		n %= 5
		n2 = n / 2

		if n100+n50+n20+n10+n5+n2 != 2 {
			fmt.Print("impossible\n")
		} else {
			fmt.Printf("possible\n")
		}
		fmt.Scan(&valuePay, &valueRecieve)
	}

}
