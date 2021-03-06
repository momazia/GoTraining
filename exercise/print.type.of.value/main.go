/*
	Source: 		Column AO in Gradebook
	Author: 		Mohamad Mahdi Ziaee
	Description:	Create a program that shows the type of some variable (use fmt.Printf)
*/
package main

import (
	"fmt"
)

func main() {
	
	// Integer
	printVariableType(12)
	
	// String
	printVariableType("Hello there!")
	
	// Float
	printVariableType(12.3)
	
	// Boolean
	printVariableType(true)
}

// Prints type of the given variable
func printVariableType(variable interface{}) {
	fmt.Printf("The variable type is: %T\n", variable)
}
