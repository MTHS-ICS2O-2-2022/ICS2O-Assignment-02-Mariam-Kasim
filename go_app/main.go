// Created by: Mariam Kasim
// Created on: March 2023
//
// This program finds the area of a square

package main

import "fmt"

func main() {
	// This function finds the area of a square
	var sidea int
	var sideb int
	var area int

	// input
	fmt.Println("This program finds the area of a square.")
	fmt.Println()
	fmt.Print("Enter side-a (cm): ")
	fmt.Scanln(&sidea)
	fmt.Print("Enter side-b (cm): ")
	fmt.Scanln(&sideb)

	// process
	area = sidea * sideb

	// output
	fmt.Println()
	fmt.Println("The area is:", area, "cm².")
	fmt.Println("\nDone.")
}
