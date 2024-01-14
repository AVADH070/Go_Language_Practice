package main

import (
	"fmt"
	"time"
)

func main() {

	go somthings_else()
	// fmt.Println(time.Second)
	time.Sleep(2 * time.Second)
	fmt.Print("This program is finished")

	go func ()  {
		fmt.Println("This is anonymous function")
	}()
	time.Sleep(1 * time.Second)

}

func somthings_else() {
	for i := 1; i <= 250000; i++ {
		fmt.Println("print number is :", i)
	}
}
