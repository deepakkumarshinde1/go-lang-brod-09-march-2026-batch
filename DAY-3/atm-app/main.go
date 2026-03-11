package main

import "fmt"

var userBalance = 100
var userPin = 1234

func exitApp(){
	fmt.Println("Thank you for visiting.")
}

func checkBalance(userBalance int){
	fmt.Println("Your balance is ₹", userBalance)
}

func cashWithdraw(){
	fmt.Println("\n Enter amount:")
	userAmount := 0
	fmt.Scan(&userAmount)

	fmt.Println("\n Enter pin:")
	pin := 0
	fmt.Scan(&pin)

	if pin == userPin{
			if userAmount > userBalance{
				fmt.Println("insufficient amount.")
			}else{
				fmt.Println("amount withdraw successful.")
				fmt.Println("your remaining amount is ₹ ",userBalance-userAmount)
			}
	}else{
		fmt.Println("Entered Pin is incorrect")
	}
}

func menu(){
	fmt.Println("\n Select and option")
	fmt.Println("1. Withdraw the cash")
	fmt.Println("2. Check balance")
	fmt.Println("3. To exit")

	choice := 0
	fmt.Scan(&choice)

	switch choice{
	case 1:
		cashWithdraw()
	case 2:
		checkBalance(userBalance)
	case 3:
		exitApp()
	default:
		fmt.Println("Invalid choice, try again")
		 menu()
	}
}

func main(){
	fmt.Println("Welcome to BOB")
	menu()
}