package module1

import "fmt"

type Person struct {
	Name         string
	Surname      string
	Relationship string
	Age          int
	Hobbies      []string
}

func AllocNewPersonPtr(
	name, surname, relationship string, age int,
	hobbies []string) *Person {
	return &Person{
		Name:         name,
		Surname:      surname,
		Relationship: relationship,
		Age:          age,
		Hobbies:      hobbies,
	}
}

func (p Person) Greet() string {
	greeting := fmt.Sprintf("Hello, my name is %s!", p.Name)
	return greeting
}

func (p Person) IsFamily() (bool, string) {
	possibleRelatives := []string{"dad", "mom", "grandpa", "grandma", "granny", "brother", "sister"}

	for _, relative := range possibleRelatives {
		if p.Relationship == relative {
			return true, "That's my family!"
		}
	}
	return false, "That's not my family!"
}

func (p *Person) AddHobby(hobby string) string {
	p.Hobbies = append(p.Hobbies, hobby)
	return formatSlice(p.Hobbies)
}
