package module1

import (
	"fmt"
)

type MyBro struct {
	Person
	Job      string
	CarBrand string
	CarModel string
}

func AllocNewMyBroPtr(
	name, surname, relationship string,
	age int, hobbies []string,
	job, carBrand, carModel string) *MyBro {
	return &MyBro{
		Person:   *AllocNewPersonPtr(name, surname, relationship, age, hobbies),
		Job:      job,
		CarBrand: carBrand,
		CarModel: carModel,
	}
}

func (p MyBro) Describe() string {
	description := fmt.Sprintf("This is my best brother ever (because I've got a few of 'em.).\nHe's %d y.o. and kinda loves %v as his hobbies.\nHe has his own %s %s.", p.Age, formatSlice(p.Hobbies), p.CarBrand, p.CarModel)
	return description
}
