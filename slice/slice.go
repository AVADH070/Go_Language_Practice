package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := [6]int{1,3,4,5,4,5}
	
	sc := arr[0:3]

	sc[0] = 10
	sort.Ints(sc)
	fmt.Println(arr)

}
