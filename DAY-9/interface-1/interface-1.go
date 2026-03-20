package main

import "fmt"

// interface majorly focus on behavior not a data
type UserInterface interface {
	printName()
}

type Student struct {
	name  string
	class string
}

func (s Student) printName() {
	fmt.Println("student name is ", s.name)
}

type Teacher struct {
	name        string
	designation string
}

func (t Teacher) printName() {
	fmt.Println("teacher name is ", t.name)
}

type Account struct {
	user UserInterface
}

func main() {
	user := Teacher{
		name: "Amit",
	}
	account := Account{
		user: user,
	}

	account.user.printName()
	// student := Student{
	// 	name: "Deepkkuamr",
	// }

	// teacher := Teacher{
	// 	name: "Amit",
	// }

	// student.printName()
	// teacher.printName()
}
