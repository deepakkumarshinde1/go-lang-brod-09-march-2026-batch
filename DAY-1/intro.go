package main // to provide a name of package

import (
	"fmt"
	// "math"
)

var name string = "Deepakkumar" // global scope

func myName() {
	text := "hi"
	fmt.Println(name, text)

	const pi = 10
	fmt.Println(pi)
}
func main() {
	//fmt.Println(text)
	myName()
	//fmt.Println("Hello to and welcome to session")
	// fmt.Println(math.Sqrt(16))

	//fmt.Println("Hello")
	//fmt.Println("Session ")

	//fmt.Printf("My age is %d \n", 36)

	// var varName dataType = value

	var v1, v2 string = "hello", "and welcome to session"

	var (
		n1        int     = 10
		isPresent bool    = true
		marks     float64 = 30.05
	)

	var age = 36 // type interface

	// short hand

	locationName := "Nashik"
	m1, m2 := 10, 20

	fmt.Println(m1, m2)

	fmt.Println(locationName)

	fmt.Println(age)

	fmt.Println(name)
	fmt.Println(v1, v2)

	fmt.Println(n1, isPresent, marks)
} // main entry  function
