package main

import "fmt"

func main(){
	result := getValues(10,20,add)
	fmt.Println(result)
}

//Passing a function as a parameter.
func getValues(a int,b int,f func(int,int) int) int{
	// result := f(a,b)
	// return result
	return f(a,b)
}

func add(a,b int) int{
	return  a+b
}

