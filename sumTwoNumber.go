package main

import "fmt"

func main() {
	var numA, numB float64

	fmt.Println("Sum Two Numbers with input from console.")
	fmt.Print("Number A : ")
	fmt.Scanf("%f", &numA)

	fmt.Print("Number B : ")
	fmt.Scanf("%f", &numB)

	result := numA + numB

	fmt.Println(numA, "+", numB, "=", result)
}
