// structs2
// Make me compile!
//
package main

import "fmt"

type ContactDetails struct {
    phone string
}

type Person struct {
	// don't just create the phone field here. embed a new struct
    ContactDetails
	name string
	age int
}

func main() {
	person := Person{name: "John", age: 32}
    person.ContactDetails = ContactDetails{phone: "+11 43 532 2135"}
	fmt.Printf("%s is %d years old and his phone is %s\n", person.name, person.age, person.phone)
}
