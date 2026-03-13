package main

import "fmt"

func main(){
	number := 10
	updateNumber(number)
	fmt.Println(number)
}

// a new copy
func updateNumber(a int){
	a = 100
}

