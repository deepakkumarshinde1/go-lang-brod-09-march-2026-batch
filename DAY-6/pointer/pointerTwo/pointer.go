package main

import "fmt"

func main(){
	number := 10;

	fmt.Println("before",number)

	var p *int
	p = &number

	*p = 100;

	fmt.Println("after",number)
	// fmt.Println(*p)
}
