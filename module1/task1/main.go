package task1

import (
	"VladOdintsov/module1/task1/models"
	"fmt"
)

type Describer interface {
	Describe() string
}

func main() {
	teacherHobbies := []string{"gym", "code", "english"}
	teacher := models.AllocNewPerson("Arthur", "Havor",
		"Teacher", 23, teacherHobbies)
	fmt.Println(teacher.Greet())
	fmt.Println(teacher.IsFamily())
	fmt.Println(teacher.AddHobby("teaching students"), "\n")

	myHobbies := []string{"cycling across the streets", "learning everything fundamentally"}
	myself := models.AllocNewMe("Dima", "Korsakov",
		"Myself", 18, myHobbies,
		"duikt", "software engineering")
	fmt.Println(myself.Describe(), "\n")

	myBroHobbies := []string{"engineering", "earning money", "wife", "physics"}
	myBro := models.AllocNewMyBro("Alexandr", "Korsakov",
		"brother", 26, myBroHobbies, "boeing",
		"ford", "focus")
	fmt.Println(myBro.Describe())
	fmt.Println(myBro.IsFamily())
	fmt.Println()

	momsHobbies := []string{"cooking", "learning english", "helping people"}
	mom := models.AllocNewRelative("Svetlana", "Korsakova",
		"mom", 47, momsHobbies, true)
	fmt.Println(mom.Greet())
	fmt.Println(mom.IsFamily())
	fmt.Println(mom.AddHobby("worrying about her sons"))
	fmt.Println()
}
