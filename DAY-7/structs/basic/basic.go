package main

import "fmt"

// struct
// type StructName struct{
// 	prop dataType
// 	prop dataType
// 	prop dataType
// 	prop dataType
// }

type Student struct{
	name string
	age int
	city string
	marks float64
	isPresent bool
}// Structure definition.

func main(){

	fmt.Println("Working...")
	// create instance
	var stud Student

	// add or update
	stud.age = 36
	stud.city = "nashik"
	stud.isPresent = true
	stud.marks = 49
	stud.name = "Deepakkumar"

	// fmt.Println(stud)
	fmt.Printf("%+v", stud)

	// print a value
	fmt.Println("\n",stud.name)
}