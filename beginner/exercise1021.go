package main

import "fmt"

func main() {
	var n float64
	fmt.Scan(&n)
	aux := int(n)
	aux2 := int(100*n) - aux*100
	fmt.Print("NOTAS:\n")
	fmt.Printf("%d nota(s) de R$ 100.00\n", int(n/100))
	aux %= 100
	fmt.Printf("%d nota(s) de R$ 50.00\n", aux/50)
	aux %= 50
	fmt.Printf("%d nota(s) de R$ 20.00\n", aux/20)
	aux %= 20
	fmt.Printf("%d nota(s) de R$ 10.00\n", aux/10)
	aux %= 10
	fmt.Printf("%d nota(s) de R$ 5.00\n", aux/5)
	aux %= 5
	fmt.Printf("%d nota(s) de R$ 2.00\n", aux/2)
	aux %= 2

	fmt.Print("MOEDAS:\n")

	fmt.Printf("%d moeda(s) de R$ 1.00\n", aux)

	fmt.Printf("%d moeda(s) de R$ 0.50\n", aux2/50)
	aux2 %= 50
	fmt.Printf("%d moeda(s) de R$ 0.25\n", aux2/25)
	aux2 %= 25
	fmt.Printf("%d moeda(s) de R$ 0.10\n", aux2/10)
	aux2 %= 10
	fmt.Printf("%d moeda(s) de R$ 0.05\n", aux2/5)
	aux2 %= 5
	fmt.Printf("%d moeda(s) de R$ 0.01\n", aux2)

}
