package main

import (
	"fmt"
	"strconv"
)

type CustomError1 struct {
	Code    int
	Message string
}

func (ce *CustomError1) Error() string {
	return fmt.Sprintf("Error Code: %d, Message: %s", ce.Code, ce.Message)
}

type CustomError2 struct {
	Message string
}

func (ce *CustomError2) Error() string {
	return fmt.Sprintf("Custom Error: %s", ce.Message)
}

func performOperation(a, b int) (int, error) {
	if b == 0 {
		return 0, &CustomError1{Code: 1001, Message: "Cannot divide by zero"}
	} else if b > 10 {
		return 0, &CustomError2{Message: "The second number should be less than or equal to 10"}
	}
	return a / b, nil
}

func main() {
	var input1, input2 string

	fmt.Print("Enter the first number: ")
	_, err := fmt.Scanln(&input1)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Print("Enter the second number: ")
	_, err = fmt.Scanln(&input2)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	num1, err := strconv.Atoi(input1)
	if err != nil {
		fmt.Println("Invalid first number. Please enter a valid number.")
		return
	}

	num2, err := strconv.Atoi(input2)
	if err != nil {
		fmt.Println("Invalid second number. Please enter a valid number.")
		return
	}

	result, err := performOperation(num1, num2)
	if err != nil {
		switch v := err.(type) {
		case *CustomError1:
			fmt.Println("Custom Error 1:", v)
		case *CustomError2:
			fmt.Println("Custom Error 2:", v)
		default:
			fmt.Println("Error:", err)
		}
		return
	}

	fmt.Println("Result:", result)
}
