package main

import "fmt"

// func funcName(?parameter){
//. function body
//}// function definition

func sub(){
		fmt.Println("I am sub function")
	}

func main(){
	// fmt.Println("Welcome to the function")
	// add()
	// sub()
	// sub()
	// sub()
	// sub()
	
	printName("Deepak")// function call we pass a data
	printName("Suraj")
	printName("Neha")
	printName("Mohan")
}

func printName(userName string){
	fmt.Println("This is my name ,", userName)
}// function definition data 
// function call
func add(){
	fmt.Println("I am add function")
}// function definition