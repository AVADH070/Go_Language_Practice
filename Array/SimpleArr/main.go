package main

import "fmt"

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	fmt.Print(arr)

	// Slicing
	arr1 := arr[1:5]
	fmt.Print(arr1)
}
