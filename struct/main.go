package main

import (
	"fmt"
)

type Person struct {
	Name      string
	Age       int
	Gender    string
	Interests []string
}

func NewPerson() Person {
	return Person{
		Name:      "John Doe",
		Age:       18,
		Gender:    "unknown",
		Interests: []string{},
	}
}

func (p *Person) AddInterest(interest string) {
	p.Interests = append(p.Interests, interest)
}

func (p *Person) HasInterest(interest string) bool {
	for _, i := range p.Interests {
		if i == interest {
			return true
		}
	}
	return false
}

func main() {
	person := NewPerson()
	person.AddInterest("coding")
	person.AddInterest("reading")
	person.AddInterest("traveling")

	fmt.Println(person.HasInterest("coding"))   // Output: true
	fmt.Println(person.HasInterest("swimming")) // Output: false
}
