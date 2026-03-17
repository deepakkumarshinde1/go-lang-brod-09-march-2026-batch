package main

import "fmt"

type Student struct{
	name string
	age int
}

func main(){
	// short hand way
	student := Student{
		name : "Deeepakkumar",
		age: 36,
	}

	fmt.Printf("%+v", student)
}
