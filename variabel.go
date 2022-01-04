package main

import "fmt"

func main() {
	// Create variable with type data
	var name string
	var age int8

	name = "Naufal Taufiq Ridwan"
	age = 19

	fmt.Println("My name is", name+". I'm", age, "years old.")  // Print Variables with println
	fmt.Printf("My name is %s. I'm %d years old.\n", name, age) // Print Variables with printf

	var firstName, lastName = "Naufal", "Taufiq Ridwan" // Create multiple varibles with value
	fmt.Println("Fullname :", firstName, lastName)

	university := "Institut Teknologi Sumatera" // Create variable without var
	println("Education :", university)
}
