package main

import "fmt"

func main(){
	getValues(10,20,add)
	getValues(10,20,sub)
}

//Passing a function as a parameter.
func getValues(a int,b int,f func(int,int)){
	f(a,b)
}

func add(a,b int){
	fmt.Println("Sum of function is ", a+b)
}

func sub(a,b int){
	fmt.Println("Sub of function is ", a-b)
}