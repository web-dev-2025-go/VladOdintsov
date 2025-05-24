package module1

import (
	"fmt"
)

type Relative struct {
	Person
	IsLoved bool
}

func AllocNewRelativePtr(name, surname, relationship string, age int, hobbies []string, isLoved bool) *Relative {
	return &Relative{
		Person:  *AllocNewPersonPtr(name, surname, relationship, age, hobbies),
		IsLoved: isLoved,
	}
}

func (p Relative) Describe() string {
	description := fmt.Sprintf("%s - my %s, he/she is %d y.o.", p.Name, p.Relationship, p.Age)
	return description
}
