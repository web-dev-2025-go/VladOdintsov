package models

import (
	"VladOdintsov/module1/task1/utils"
	"fmt"
)

type Person struct {
	Name         string
	Surname      string
	Relationship string
	Age          int
	Hobbies      []string
}

func AllocNewPerson(
	name, surname, relationship string, age int,
	hobbies []string) Person {
	return Person{
		Name:         name,
		Surname:      surname,
		Relationship: relationship,
		Age:          age,
		Hobbies:      hobbies,
	}
}
func (p Person) Greet() string {
	greeting := fmt.Sprintf(
		"Hello, my name is %s!", p.Name)
	return greeting
}
func (p Person) IsFamily() (bool, string) {
	possibleRelatives := []string{
		"dad", "mom", "grandpa", "grandma", "granny", "brother", "sister"}
	for i := range possibleRelatives {
		if p.Relationship == possibleRelatives[i] {
			return true, "That's my family!"
		}
	}
	return false, "That's not my family!"
}
func (p *Person) AddHobby(hobby string) string {
	p.Hobbies = append(p.Hobbies, hobby)
	return utils.FormatSlice(p.Hobbies)
}

type Me struct {
	Person
	University string
	MajorSubj  string
}

func AllocNewMe(
	name, surname, relationship string,
	age int, hobbies []string,
	university, majorSubj string) Me {
	return Me{
		Person:     AllocNewPerson(name, surname, relationship, age, hobbies),
		University: university,
		MajorSubj:  majorSubj,
	}
}
func (p Me) Describe() string {
	return fmt.Sprintf(
		"My name is %s %s. I'm %d years old and I've got %v as my favourite stuff.\n"+
			"Studying %s in %s",
		p.Name, p.Surname, p.Age, utils.FormatSlice(p.Hobbies), p.MajorSubj, p.University)
}

type MyBro struct {
	Person
	Job      string
	CarBrand string
	CarModel string
}

func AllocNewMyBro(
	name, surname, relationship string,
	age int, hobbies []string,
	job, carBrand, carModel string) MyBro {
	return MyBro{
		Person:   AllocNewPerson(name, surname, relationship, age, hobbies),
		Job:      job,
		CarBrand: carBrand,
		CarModel: carModel,
	}
}
func (p MyBro) Describe() string {
	return fmt.Sprintf(
		"This is my best brother ever (because I've got a few of 'em.).\n"+
			"He's %d y.o. and kinda loves %v as his hobbies.\n"+
			"He has his own %s %s.",
		p.Age, utils.FormatSlice(p.Hobbies), p.CarBrand, p.CarModel)
}

type Relative struct {
	Person
	IsLoved bool
}

func AllocNewRelative(
	name, surname, relationship string,
	age int, hobbies []string, isLoved bool) Relative {
	return Relative{
		Person:  AllocNewPerson(name, surname, relationship, age, hobbies),
		IsLoved: isLoved,
	}
}
func (p Relative) Describe() string {
	return fmt.Sprintf(
		"%s - my %s, he/she is %d y.o.",
		p.Name, p.Relationship, p.Age)
}
