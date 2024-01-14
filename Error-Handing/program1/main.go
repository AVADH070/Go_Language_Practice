package main

import (
	"fmt"
)

type CustomError struct {
	Code    int
	Message string
}

type inputError struct {
	Message string
}

func (ce CustomError) Error() string {
	return fmt.Sprintf("Error Code: %d, Message: %s", ce.Code, ce.Message)
}

func (ie inputError) Error() string {
	return fmt.Sprintf("Please Type number not a  %v", ie.Message)
}

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, CustomError{Code: 1001, Message: "Cannot divide by zero"}
	}
	return a / b, nil
}
func inputChacked(input string) (int, error) {
	return 45, nil
}

func main() {

	result, err := divide(10, 0)
	if err != nil {
		customErr, ok := err.(CustomError)
		fmt.Printf("%T ", customErr)
		if ok {
			fmt.Println(customErr)
		} else {
			fmt.Println("Error:", err)
		}
	} else {
		fmt.Println("Result:", result)
	}
}
