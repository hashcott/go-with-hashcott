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
	fmt.Println(hanh)
}
