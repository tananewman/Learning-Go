package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Hawkins",
		contactInfo: contactInfo{
			email:   "12354@spankme.net",
			zipCode: 84062,
		},
	}

	jim.print()
	jim.updateFirstName("Dickhole")
	jim.print()
}

func updateTestString(strPtr *string, newVal string) {
	*strPtr = newVal
}

func (personPtr *person) updateFirstName(newFirstName string) {
	(*personPtr).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
