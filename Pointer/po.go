package main

import "fmt"

func main() {
	x := 10
	y := &x

	*y = 50
	fmt.Println(x) // Change the value of x
	// fmt.Printf("%T", y) // *int

	// var slice = []int{1, 2, 3, 4, 4, 5, 5}
	// fmt.Println(slice[4:])
}
