package main

import "fmt"

// 1st Way :: Function written as a definition without a parameter.


func mainFirstWay(){
	getValueOne(10,30)
}

func getValueOne(a int,b int) func() {
	return func() {
		fmt.Println("final sum is ",a+b)
	}// definition
}

// 2. Way :: Function return definition with return type.
func mainSecondWay(){
	add := getValueTwo(10,30)
	add()
}
func getValueTwo(a int,b int) func(  ) int{
	return func() int{
		fmt.Println("final sum is ",a+b)
	}// definition
}

func main(){
	add := getValue(10,30)
	result := add(100)
	fmt.Println("final sum is ",result)
}
// Returning of function definition.

// 3. Way :: function written definition with parameter and return type.
func getValue(a int,b int) func( int ) int{
	return func(c int) int{
		return a+b+c // 10+30+100
	}// definition
}