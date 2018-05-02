package main

import "fmt"

func main() {
	var number int
	var ans float64
	fmt.Scan(&number)

	for i := 0; i < number; i++ {
		ans += 2.0
		ans = 1.0 / ans
	}
	ans += 1.0

	fmt.Printf("%.10f\n", ans)
}
