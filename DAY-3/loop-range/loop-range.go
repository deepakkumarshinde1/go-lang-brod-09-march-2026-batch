package main

import "fmt"

func main(){
	// number := 10
	text := "Deepakkumar"
	
	for i,value := range text{
		fmt.Println(i, string(value))
	}
}