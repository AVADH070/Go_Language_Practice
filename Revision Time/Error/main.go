package main

import (
	"fmt"
	"strconv"
)

type inputError11 struct {
	Message string
}

func (i *inputError11) Error() string {
	return fmt.Sprintf("%s ", i.Message)
}

func chackedNumber(n string) (int, error) {
	_, err := strconv.Atoi(n)
	if err != nil {
		return 0, &inputError11{Message: "Please enter valid number"}
	} else {
		fmt.Print("The number is : ", n)
	}
	return 1, nil
}

func main() {
	var num string
	fmt.Print("Enter a number : ")
	fmt.Scanln(&num)

	_, err := chackedNumber(num)
	if err != nil {
		fmt.Print(err)
	}

}
