// Created by: Janet Do
// Created on: Sep 2020
//
// This program displays a user's address
package main

import "fmt"

func main() {
	// This function displays your address when typed in
	var streetName string
	var streetNumber int

	// input
	fmt.Println("This program displays a user's address")
	fmt.Println()
	fmt.Print("Enter your street number: ")
	fmt.Scanln(&streetNumber)
	fmt.Print("Enter your street name: ")
	fmt.Scanln(&streetName)

	// output
	fmt.Println("Your Address is:", streetNumber, streetName, ".")

	fmt.Println("\nDone.")
}
