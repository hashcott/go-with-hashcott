package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}
type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	// hanh := person{firstName: "Nguyen", lastName: "Hanh"}
	// var hanh person

	// fmt.Println(hanh)
	// fmt.Printf("%+v\n", hanh)
	// hanh.firstName = "Nguyen"
	// hanh.lastName = "Hanh"
	// hanh.contact = contactInfo{
	// 	email:   "Hanh@gmail.com",
	// 	zipCode: 2536,
	// }
	hanh := person{
		firstName: "Nguyen",
		lastName:  "Hanh",
		contact: contactInfo{
			email:   "hanh@gmail.com",
			zipCode: 2535,
		},
	}
	hanhPointer := &hanh
	hanhPointer.updateName("John")
	// hanh.print()

	a := 5
	test(&a)
	fmt.Println(a)
}

func test(a *int) {
	*a = 10
}
func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
