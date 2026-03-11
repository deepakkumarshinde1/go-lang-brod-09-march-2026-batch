package main

import "fmt"

func main(){

	fmt.Println("--- Menu ----")
	fmt.Println("1. Java")
	fmt.Println("2. Go Lang")
	fmt.Println("3. Node JS")
	fmt.Println("4. Kotlin")

	fmt.Println("Enter Course Number")
	courseNumber := 0
	fmt.Scan(&courseNumber)

	switch courseNumber{
	case 1:
		fmt.Println("Congratulations you have selected Java Course")
	case 2:
		fmt.Println("Congratulations you have selected Go Lang Course")
	case 3:
		fmt.Println("Congratulations you have selected Node JS Course")
	case 4:
		fmt.Println("Congratulations you have selected Kotlin Course")
	default:
		fmt.Println("Invalid Choice")
	}


}