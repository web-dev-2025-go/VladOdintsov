package main

import (
	"VladOdintsov/module1"
	"fmt"
)

func main() {
	fmt.Println("This line is written in PowerShell\n")

	var teacherHobbies []string = []string{"gym", "code", "english"}
	var teacher module1.Person = *module1.AllocNewPersonPtr("Arthur", "Havor",
		"Teacher", 23, teacherHobbies)
	fmt.Println(teacher.Greet())
	fmt.Println(teacher.IsFamily())
	fmt.Println(teacher.AddHobby("teaching students"), "\n") // я дуже здивувався коли вбудована println() не вивела слайс, а вивела адресу в пам'яті, але це стало мені уроком

	myHobbies := []string{"cycling across the streets", "learning everything fundamentally"}
	myself := *module1.AllocNewMePtr("Dima", "Korsakov",
		"Myself", 18, myHobbies,
		"duikt", "software engineering")
	fmt.Println(myself.Describe(), "\n")

	var myBroHobbies = []string{"engineering", "earning money", "wife", "physics"}
	var myBro = *module1.AllocNewMyBroPtr("Alexandr", "Korsakov",
		"brother", 26, myBroHobbies, "boeing",
		"ford", "focus")
	fmt.Println(myBro.Describe())
	fmt.Println(myBro.IsFamily())
	fmt.Print("\n")

	var momsHobbies []string = make([]string, 3)
	momsHobbies[0] = "cooking"
	momsHobbies[1] = "learning english"
	momsHobbies[2] = "helping people"
	mom := *module1.AllocNewRelativePtr("Svetlana", "Korsakova",
		"mom", 47, momsHobbies, true)
	fmt.Println(mom.Greet())
	fmt.Println(mom.IsFamily())
	fmt.Println(mom.AddHobby("worrying about her sons"))
	fmt.Print("\n")
}
