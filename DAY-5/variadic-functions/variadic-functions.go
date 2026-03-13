package main

import "fmt"

func main(){
	numbers := []int{10,20,1}
	result := add(numbers...)
	fmt.Println(result)

}

func add(a ...int) int{
	result := 0

	for _,value := range a{
		result += value
	}

	return result
}