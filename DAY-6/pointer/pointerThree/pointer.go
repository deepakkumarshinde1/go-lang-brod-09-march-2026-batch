package main

import "fmt"

// pass a pointer in Function parameter
func main(){
	number := 100
	fmt.Println("outer before",number)
	increment(&number)
	increment(&number)
	increment(&number)
	fmt.Println("outer after",number)
}

func increment(num *int){
	*num++
	// fmt.Println("inner",*num)
}


// So function as a value as parameter.
// Each time there will be a new copy for num.
// func increment(num int){
// 	num++
// 	fmt.Println("inner",num)
// }

