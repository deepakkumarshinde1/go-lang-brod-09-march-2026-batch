package main

import "fmt"

func main(){
	
	number := 10
	// var p *int
	p := &number
	pp := &p
	ppp := &pp
	fmt.Println("p",*p)
	fmt.Println("pp",**pp)
	fmt.Println("pp",***ppp)
}
