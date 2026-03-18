package main

import "fmt"

type Student struct{
	name string
	class string
	marks float64
}
// Okay, so we have a concept that is the receiver.
//Receiver is a syntax which is written in between function keyword and a function name.
func (s Student) updateName(){
		fmt.Printf("%+v \n",s)
}

func main(){
	// So to create a struct instance
	student := Student{
		name: "Deepak",
		class: "12th",
		marks: 303,
	}

	student1 := Student{
		name:"Sam",
		class: "11th",
	}
	student.updateName()
	student1.updateName()
	// student.name = "Arun"

	// fmt.Printf("%+v \n",student.name)
}