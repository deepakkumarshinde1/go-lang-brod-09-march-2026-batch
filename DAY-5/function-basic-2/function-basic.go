package main

import "fmt"

func main(){
	add(10,20)
	result :=sub(10,5)
	fmt.Println("Sub result ",result)
}

// how to pass multiple value to a function
// func add(a int,b int){
// 	fmt.Println("Sum of a & b : ",a+b)
// }

func add(a ,b int){
	fmt.Println("Sum of a & b : ",a+b)
}


// function with return value

// func Name(para) returnType{}
func sub(a,b int) int {
	result := a-b
	return result
}


