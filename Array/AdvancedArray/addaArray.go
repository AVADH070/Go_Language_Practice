package main

import "fmt"

func main() {
	// Arrays
	arr1 := [3]int{9, 7, 6}
	arr2 := []int{9, 7, 1}
	arr3 := [3]int{9, 5, 3}

	// fmt.Println(arr1[:3] == arr2)
	// fmt.Println(arr2 == arr3)
	// fmt.Println(arr1 == arr3)

	fmt.Println(arr1 == arr3)
	fmt.Println(arr2 == arr3[:3]) // Comparing the first 3 elements of arr3 with arr2

}
