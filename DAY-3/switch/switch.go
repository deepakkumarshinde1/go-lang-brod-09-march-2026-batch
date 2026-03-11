package main

import "fmt"

func main(){
	fmt.Println("Working ...")

	number := 110

	switch number{
	case 10:
		fmt.Println("Number is 10")
	case 50:
		fmt.Println("Number is 50")
	case 100:
		fmt.Println("Number is 100")
	default:
		fmt.Println("Number is no equal ")
	}

	// number = 17
	// switch number{
	// case 1,3,5,7,9:
	// 	fmt.Println("number is odd")

	// case 0,2,4,6,8,10:
	// 	fmt.Println("Number is even")

	// default:
	// 	fmt.Println("Please enter a number less then or equals to 10")
	// }

	
	switch number := 17;number{
	case 1,3,5,7,9:
		fmt.Println("number is odd")

	case 0,2,4,6,8,10:
		fmt.Println("Number is even")

	default:
		fmt.Println("Please enter a number less then or equals to 10")
	}

	switch num:= 1;{
	case num > 5:
		fmt.Println("Number is greater the 5")
	case num > 2:
		fmt.Println("Number is greater the 2 but less 5")
	default :
		fmt.Println("Number is less then or equals to  2")
	}
}