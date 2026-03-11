package main

import "fmt"

func main(){
	fmt.Println("Welcome Admin")
	adminPassword := 1234
	i := 0
	for i=1; i <= 3 ;i++{
		fmt.Println("Enter you password")
		var password int

		fmt.Scan(&password)

		if password == adminPassword{
			fmt.Println("Login Successfully")
			break;
		}else if i < 3{
			fmt.Println("Wrong Password , try again")
		}
	}

	if i == 4{
		fmt.Println("Your account is blocked for 1hr")
	}
}