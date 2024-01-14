package main

import "fmt"

func main() {
	// var ch_name chan int

	// ch_name = make(chan int) // ch_name := make(chan int)
	ch_name := make(chan int)

	go func (ch chan int)  {
		fmt.Println(<-ch)
	}(ch_name)

	ch_name <- 10
	fmt.Print(<-ch_name)
}
