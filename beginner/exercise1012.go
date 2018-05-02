package main

import (
	"fmt"
)

func main() {
	const pi = 3.14159
	//base | base maior | altura | area triangulo | area do circulo | area do trapezio | area do quadrado | area do retangulo
	var base, largeBase, height float64

	fmt.Scan(&base, &largeBase, &height)

	triangleArea := (base * height) / 2
	circleArea := pi * height * height
	trapezoidArea := ((base + largeBase) * height) / 2
	squareArea := largeBase * largeBase
	rectangleArea := base * largeBase

	fmt.Printf("TRIANGULO: %.3f\n", triangleArea)
	fmt.Printf("CIRCULO: %.3f\n", circleArea)
	fmt.Printf("TRAPEZIO: %.3f\n", trapezoidArea)
	fmt.Printf("QUADRADO: %.3f\n", squareArea)
	fmt.Printf("RETANGULO: %.3f\n", rectangleArea)
}
