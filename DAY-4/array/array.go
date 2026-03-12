package main

import "fmt"

func main(){
	fmt.Println("Working")

	// var array [size]dataType
	// var array = [size]dataType{}

	// var array [5]int
	// var array [5]string
	// var array [5]float64

	// int => 0
	// string => ""
	// bool => false
	// float => 0

	// 0 => 10 , 1 => 20
	var array = [5]int{10,20,30,40,50}
	
	// array[4] = 30


	// var array1 = [5]int{10,20}
	// fmt.Println(array == array1)

	// single => using index
	// fmt.Println(array[0])

	// multiple => loop
	// fmt.Println(len(array))

	// for i := 0; i < len(array); i++{
	// 	fmt.Println(array[i])
	// }

	for i := len(array)-1; i >= 0; i--{
		fmt.Println(array[i])
	}

	// for i,value:=range array{
	// 	fmt.Println(i,value)
	// }

	// fmt.Println(array[4])

	number := [...]int{1,2,3,4}
	fmt.Println(len(number))

}