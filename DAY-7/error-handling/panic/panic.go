package main

import "fmt"

func main(){
	var p *int

	defer func(){
		if r := recover(); r != nil{
			fmt.Println(r)
		}
	}()

	if(p == nil){
		panic("Pointer p is null , please assign a value to it")
	}
	fmt.Println(p)
}
