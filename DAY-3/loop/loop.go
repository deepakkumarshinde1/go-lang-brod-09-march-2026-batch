package main

import "fmt"


func main(){
	// for 
	for number := 1 ; number <= 5 ; number++{
		fmt.Println("for", number)
	}

	// like while
	start := 0
	end := 5

	for start <= end{
		fmt.Println("while ", start)
		start++
	}

	// break => stop
	for i:=1 ; i < 10 ; i++{
		if(i % 5 == 0){
			break
		}

		fmt.Println(i)
	}
	// continue => skip and the start

	for i:=1 ; i < 10 ; i++{
		if(i % 2 == 0){
			continue
		}
		fmt.Println(i)
	}

	// for infinity
	for{
		fmt.Println("loop")
	}

	


}