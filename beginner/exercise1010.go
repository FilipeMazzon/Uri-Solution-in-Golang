package main

import (
	"fmt"
)

func main() {
	var pieceCode1, PieceNumber, pieceCode2, PieceNumber2 int
	var vup1, vup2 float64

	fmt.Scan(&pieceCode1, &PieceNumber, &vup1, &pieceCode2, &PieceNumber2, &vup2)

	valuePay := float64(PieceNumber) * vup1
	valuePay2 := float64(PieceNumber2) * vup2

	allValue := valuePay + valuePay2

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", allValue)

}
