package main

import "fmt"

// anonymous functions.
func main(){
	add := func (a,b int){
		fmt.Println("Hello" , a+b)
	};

	// addOne := add
	add(10,20)
	// addOne(30,40)
}

