package main

import (
	// "booking-app/helper"
	"fmt"
	"sync"
	"time"
	// "strconv"
)

// package level Variable
var festName = "Cloud Fest"

const Tickets uint = 69

var BookingsList = make([]UserData, 0)
var remainingTickets uint = 69

type UserData struct {
	firstName  string
	lastName   string
	email      string
	tickets    uint
	newsLetter bool
}

var wg = sync.WaitGroup{}

func main() {

	// here we can use := instead of var
	greet()

	// fmt.Printf("1 %T , 2 %T \n", festName, Tickets)

	// for {
	firstName, lastName, email, num := getUserInput()
	isValidName, isValidEmail, validTickets := Valid(firstName, lastName, email, num, remainingTickets)

	if isValidName && isValidEmail && validTickets {
		remainingTickets -= num

		var userData = UserData{
			firstName:  firstName,
			lastName:   lastName,
			email:      email,
			tickets:    num,
			newsLetter: true,
		}
		// userData["firstName"] = firstName
		// userData["lastName"] = lastName
		// userData["email"] = email
		// userData["tickets"] = strconv.FormatUint(uint64(num), 10)

		fmt.Println("userData")
		fmt.Println(userData)

		// var BookingsList = [50]string{}
		// println(firstName, remainingTickets, festName)
		fmt.Printf("thank you %v  %v for booking tickets the details will be sent to your email %v ", firstName, lastName, email)
		fmt.Printf("remaining tickets are %v \n", remainingTickets)
		// BookingsList[0] = firstName
		// BookingsList[1] = lastName
		BookingsList = append(BookingsList, userData)
		fmt.Println("list of booking")
		fmt.Println(BookingsList)

		// for _, fnv := range BookingsList {
		// 	fmt.Println("yo yo")
		// 	fmt.Println(fnv)
		// }

		// fmt.Printf("first name list %v \n", firstN())
		// firstNamer := []string{}
		// var i = "2"
		// firstNamer = append(firstNamer, firstName)
		// fmt.Printf("%v ,dfhbv, %v", i, firstNamer)

		fmt.Printf("first name list %v \n", firstN())

		wg.Add(1)
		go sendTickets(firstName, lastName, num, email)
		// here go serprates this line of code so that it doesn't hinder the code and it's running
		// noTicket := remainingTickets == 0
		fmt.Printf("Out of %v we have %v left\n", Tickets, remainingTickets)

		// if remainingTickets == 0 {
		// fmt.Println("our seats are full come back next year")
		// break
		// }
	} else {
		if !isValidName {
			fmt.Println("you have entered an invalid name")
		}
		if !isValidEmail {
			fmt.Println("you have entered an invalid email address")
		}
		if !validTickets {
			fmt.Printf("invalid number of tickets we have %v tickets left\n", remainingTickets)
		}

	}

	// }
	wg.Wait()

}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var num uint
	fmt.Println(`enter your first name:`)
	fmt.Scan(&firstName)

	fmt.Println(`enter your last name:`)
	fmt.Scan(&lastName)

	fmt.Println(`enter your email`)
	fmt.Scan(&email)

	fmt.Printf("number of tickets:")
	fmt.Scan(&num)

	return firstName, lastName, email, num
}

func greet() {
	fmt.Printf("Hello and Welcome to %v\n", festName)
	fmt.Printf("GET your Tickets here to attend the %v here \n", festName)
	fmt.Println(`Bookings open now only`)
}

func firstN() []string {
	firstNames := []string{}
	for _, fnameV := range BookingsList {
		firstNames = append(firstNames, fnameV.firstName)
	}
	// var names = strings.Fields(fnameV)
	// var fnames = names[0]
	// fmt.Println(fnameV)
	return firstNames
}

func sendTickets(firstName string, lastName string, num uint, email string) {
	time.Sleep(5 * time.Second)
	var data = fmt.Sprintf("dear %v %v we are sending %v tickets to your email", firstName, lastName, num)
	fmt.Println("######")
	fmt.Printf("%v %v \nthank you\n", data, email)
	fmt.Println("######")
	wg.Done()
}

// example
// You can edit this code!
// Click here and start typing.
// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	fmt.Println("Hello, madarchod")
// 	str := "hello "
// 	s := strings.Fields(str)
// 	fmt.Println(s)
// 	fmt.Println(len(s))
// }
