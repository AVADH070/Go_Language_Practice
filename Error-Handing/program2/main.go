package main

import (
	"fmt"
	"strconv"
) 

// CustomError is a custom error type that implements the error interface.
type CustomError struct {
	Message string
}

// Error implements the error interface for CustomError.
func (ce CustomError) Error() string {
	return fmt.Sprintf("Custom Error: %s", ce.Message)
}

func main() {
	var input string
	fmt.Print("Enter a number: ")
	_, err := fmt.Scanln(&input)

	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Attempt to convert the input to an integer
	number, err := strconv.Atoi(input)
	if err != nil {
		// Create and handle a custom error
		customErr := CustomError{Message: "Invalid input. Please enter a valid number."}
		fmt.Println(customErr)
		return
	}

	fmt.Println("You entered:", number)
	// Do something with the number...
}
