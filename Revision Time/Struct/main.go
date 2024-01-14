package main

import "fmt"

type sampleS struct {
	name   string
	num    int
	height float64
}



// Interface

type BMIWeight interface {
	getBMIHeight() float64
}

func (s sampleS) getBMIHeight() float64 {
	return float64(s.num) * s.height
}

func calculate(b BMIWeight) {

	s, ok := b.(sampleS)
	fmt.Printf("%T\n", s)
	fmt.Printf("%v", ok)
	fmt.Print(b.getBMIHeight())
	// fmt.Print(s.getBMIHeight()) as Well
}
func main() {
	res := sampleS{
		name:   "Avadh Sonagara",
		num:    43,
		height: 32.5,
	}

	calculate(res)
}
