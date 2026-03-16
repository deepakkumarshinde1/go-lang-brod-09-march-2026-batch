package main

import "fmt"

// anonymous functions.
func main(){

	//Immediately invoke function (IIF)
	defer func (a,b int){
		fmt.Println("IIF" , a+b)
	}(30,70)

	// Anonymous function.
	add := func (a,b int){
		fmt.Println("Add AF" , a+b)
	}

	sub := func (a,b int){
		fmt.Println("Sub AF" , a-b)
	}

	defer add(10,20)
	sub(10,5)
}

