package main

import "fmt"

type Location struct{
	city string
	pin_code int
	landmark string
}

type Student struct{
	name string
	class string
	marks float64
	location Location
}

// Receiver
// written In between function keyword and function.
// s => Receiver , Student => Structure (struct)
// r as value 
// r as pointer
func (s *Student) printData(){
	s.name = "Deepak"	
	s.location.landmark = "Dwarka Circle"
}

func main(){
		student  := Student{
			name: "Suraj",
			class: "10th",
			marks: 45,
			location: Location{
				city: "Nashik",
				pin_code: 240011,
				landmark: "Near New-Palace",
			},	
		}

		student.printData()
		fmt.Printf("%+v \n", student)
	
}