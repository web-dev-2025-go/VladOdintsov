package module1

import (
	"fmt"
)

type Me struct {
	Person
	University string
	MajorSubj  string
}

func AllocNewMePtr(
	name, surname, relationship string,
	age int, hobbies []string,
	university, majorSubj string) *Me {
	return &Me{
		Person:     *AllocNewPersonPtr(name, surname, relationship, age, hobbies),
		University: university,
		MajorSubj:  majorSubj,
	}
}

func (p Me) Describe() string {
	description := fmt.Sprintf("My name is %s %s. I'm %d years old and I've got %v as my favourite stuff.\nStudying %s in %s", p.Name, p.Surname, p.Age, formatSlice(p.Hobbies), p.MajorSubj, p.University)
	return description
}
