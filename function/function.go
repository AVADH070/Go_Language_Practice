// package main

// import "fmt"

// func incr(p int, q int) (x int, y int) {
// 	x = p
// 	y = q
// 	return
// }

// func main() {
// 	f, l := incr(1, 1342)
// 	fmt.Print(f,l)
// }

// Struct Example
// ==========================================================================================

/*
type simpleStructure struct {
	name   string
	number int
	ab     Wheel
}

type Wheel struct {
	fname string
	lname string
}

func main() {
	acc := simpleStructure{
		name:   "A",
		number: 10,
		ab: Wheel{
			fname: "Avadh",
			lname: "Sonagara",
		},
	}

	fmt.Println(acc.name)
	fmt.Println(acc.number)
	fmt.Println(acc.ab.fname)
	fmt.Println(acc.ab.lname)

}
*/

// ==========================================================================================

/*
type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	// ?
	return true
}

// don't touch below this line

func test(mToSend messageToSend) {
	fmt.Printf(`sending "%s" from %s (%v) to %s (%v)...`,
		mToSend.message,
		mToSend.sender.name,
		mToSend.sender.number,
		mToSend.recipient.name,
		mToSend.recipient.number,
	)
	fmt.Println()
	if canSendMessage(mToSend) {
		fmt.Println("...sent!")
	} else {
		fmt.Println("...can't send message")
	}
	fmt.Println("====================================")
}

func main() {
	test(messageToSend{
		message: "you have an appointment tommorow",
		sender: user{
			name:   "Brenda Halafax",
			number: 16545550987,
		},
		recipient: user{
			name:   "Sally Sue",
			number: 19035558973,
		},
	})
	test(messageToSend{
		message: "you have an event tommorow",
		sender: user{
			number: 16545550987,
		},
		recipient: user{
			name:   "Suzie Sall",
			number: 0,
		},
	})
	test(messageToSend{
		message: "you have an party tommorow",
		sender: user{
			name:   "Njorn Halafax",
			number: 16545550987,
		},
		recipient: user{
			name:   "Sally Sue",
			number: 19035558973,
		},
	})
	test(messageToSend{
		message: "you have a birthday tommorow",
		sender: user{
			name:   "Eli Halafax",
			number: 0,
		},
		recipient: user{
			name:   "Whitaker Sue",
			number: 19035558973,
		},
	})
}

*/

// =========================================================================
/*

// Inheritance

type student struct {
	name string
	role string
}

type school struct {
	school string
	number string
	student
}

func main() {
	temp := school{
		school: "JNB",
		number: "34234",
		student: student{
			name: "Avadh",
			role: "Engineering",
		},
	}

	fmt.Println(temp.school)
	fmt.Println(temp.number)
	fmt.Println(temp.name)
	fmt.Println(temp.role)
}
*/

/*
package main

import (
	"fmt"
	"time"
)

// func sendMessage(msg message) {
// 	fmt.Println(msg.getMessage())
// }

type message interface {
	getMessage() string
}

// don't edit below this line

type birthdayMessage struct {
	birthdayTime  time.Time
	recipientName string
}

func (bm birthdayMessage) getMessage() string {
	return fmt.Sprintf("Hi %s, it is your birthday on %s", bm.recipientName, bm.birthdayTime.Format(time.RFC3339))
}

type sendingReport struct {
	reportName    string
	numberOfSends int
}

func (sr sendingReport) getMessage() string {
	return fmt.Sprintf(`Your "%s" report is ready. You've sent %v messages.`, sr.reportName, sr.numberOfSends)
}

func test(m message) {
	fmt.Println(m.getMessage())
	fmt.Println("====================================\n")
}

func main() {
	test(sendingReport{
		reportName:    "First Report",
		numberOfSends: 10,
	})
	test(birthdayMessage{
		recipientName: "John Doe",
		birthdayTime:  time.Date(1994, 03, 21, 0, 0, 0, 0, time.UTC),
	})
	test(sendingReport{
		reportName:    "First Report",
		numberOfSends: 10,
	})
	test(birthdayMessage{
		recipientName: "Bill Deer",
		birthdayTime:  time.Date(1934, 05, 01, 0, 0, 0, 0, time.UTC),
	})
}

*/

/*
package main

import (
	"fmt"
)

type shape interface {
	area() float64
}



type circle struct {
	redius float64
}

type ractangel struct {
	height float64
	width  float64
}

func (c circle) area() float64 {
	return 3.14 * c.redius * c.redius
}

func (r ractangel) area() float64 {
	return 3.14 * r.height * r.width
}

func (r ractangel) getHeight() float64 {
	return r.height
}

func Calculates(s shape) {
	c, ok1 := s.(circle)
	ract, ok := s.(ractangel)

	fmt.Printf("%T and %v\n",ract,ok)
	fmt.Printf("%T and %v\n",c,ok1)

	if ok1 {
		fmt.Println("The Height of ractangle is :",ract.getHeight())
		fmt.Printf("The area of ractangle is %.3f",s.area())
	}else{
		fmt.Println("It Can't Ractangle")
	}
}

func main() {
	c := circle{
		redius: 43.3,
	}

	Calculates(c)

	// r := ractangel{
	// 	height: 2.2,
	// 	width:  2.4,
	// }
	// Calculates(r)
	// r.getHeight()
}
*/

/*
package main

import (
	"fmt"
)

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

type expense interface {
	cost() float64
}

// don't touch below this line

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}

func estimateYearlyCost(e expense, averageMessagesPerYear int) float64 {
	return e.cost() * float64(averageMessagesPerYear)
}

func getExpenseReport(e expense) (string, float64) {
	// _, ok := e.(email)
	// if ok {
	// 	return "A", e.cost()
	// }
	// _, ok1 := e.(sms)

	// if ok1 {
	// 	return "B", e.cost() // Aa pan chale
	// }
	// return "", 0.0


	em, ok := e.(email)
	if ok {
		return "A", em.cost()
	}
	s, ok1 := e.(sms)

	if ok1 {
		return "B", s.cost()
	}
	return "", 0.0
}

func test(e expense) {
	address, cost := getExpenseReport(e)
	
	switch e.(type) {
	case email:
		fmt.Printf("Report: The email going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	case sms:
		fmt.Printf("Report: The sms going to %s will cost: %.2f\n", address, cost)
		fmt.Println("====================================")
	default:
		fmt.Println("Report: Invalid expense")
		fmt.Println("====================================")
	}
}

func main() {
	test(email{
		isSubscribed: true,
		body:         "hello there",
		toAddress:    "john@does.com",
	})
	test(email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
		toAddress:    "jane@doe.com",
	})
	test(email{
		isSubscribed: false,
		body:         "This meeting could have been an email",
		toAddress:    "elon@doe.com",
	})
	test(sms{
		isSubscribed:  false,
		body:          "This meeting could have been an email",
		toPhoneNumber: "+155555509832",
	})
	test(sms{
		isSubscribed:  false,
		body:          "This meeting could have been an email",
		toPhoneNumber: "+155555504444",
	})
	test(invalid{})
}

*/

package main

import "fmt"

type DivideByZeroError struct {
	dividend int
}

func (e DivideByZeroError) Error() string {
	return fmt.Sprintf("cannot divide %d by 0", e.dividend)
}


func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, DivideByZeroError{dividend: a}
	}
	return a / b, nil
}

func main() {
	_, err := Divide(2, 0)

	if err != nil {
		fmt.Println(err)
	}
}
