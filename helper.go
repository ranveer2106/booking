package main

import "strings"

func Valid(fname string, lname string, email string, num uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(fname) > 2 && len(lname) > 2
	isValidEmail := strings.Contains(email, "@")
	validTickets := num <= remainingTickets && num > 0
	return isValidName, isValidEmail, validTickets
}
