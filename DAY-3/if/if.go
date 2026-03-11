package main

import "fmt"

func main(){
	fmt.Println("Its working")
	number := 1
	if number > 2{
		fmt.Println("Yes the number is greater then 2")
	}


	if number > 2{
		fmt.Println("Yes the number is greater then 2")
	}else{
		fmt.Println("Number is less then 2")
	}

	number = 1
	if number > 5{
		fmt.Println("Yes the number is greater then 5")
	}else if number > 2{
		fmt.Println("Yes the number is greater then 2 but less the 5")
	}else{
		fmt.Println("Yes the number is less then 2")
	}

	// number = 9
	// if number > 5{
	// 	if number%2 == 0{
	// 		fmt.Printf("Number %d is even" ,number)
	// 	}else{
	// 		fmt.Printf("Number %d is odd",number)
	// 	}
	// }
	
	
	if num := 2;num > 5{
		if num%2 == 0{
			fmt.Printf("Number %d is even" ,num)
		}else{
			fmt.Printf("Number %d is odd",num)
		}
	}else{
		fmt.Printf("Number %d is less then 5",num)
	}
}