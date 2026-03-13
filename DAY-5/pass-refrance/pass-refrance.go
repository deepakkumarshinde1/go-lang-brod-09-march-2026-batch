package main

import "fmt"

func main(){
	number := 10
	fmt.Println("Before " ,number)
	updateNumber(&number)
	fmt.Println("After ", number)
}

// pass a reference
// & => get location of variable
// * => get a value of location
func updateNumber(a *int){
	*a = 100
}

