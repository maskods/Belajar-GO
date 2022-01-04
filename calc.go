package main

import (
	"fmt"
	"os"
)

func main() {
	var numA, numB, result float64
	var operation string

	fmt.Println("Simple Calculator")
	fmt.Print("Number A : ")
	fmt.Scanf("%f", &numA)

	fmt.Print("Number B : ")
	fmt.Scanf("%f", &numB)

	fmt.Print("Operation (+, -, /, *): ")
	fmt.Scanf("%s", &operation)

	switch operation {
	case "+":
		result = numA + numB
	case "-":
		result = numA - numB
	case "/":
		result = numA / numB
	case "*":
		result = numA * numB
	default:
		fmt.Println("Invalid operation")
		os.Exit(1)
	}

	fmt.Println(result)
}
