package main

import (
	"fmt"
)

func main(){
	var balance,withdrawAmount int
	
	fmt.Println("Enter you balance")
	fmt.Scan(&balance)

	fmt.Println("Enter Amount to withdraw")
	fmt.Scan(&withdrawAmount)

	message,err:= withdraw(&balance, withdrawAmount)

	if(err != nil){
		fmt.Println("Error:",err)
	}else{
		fmt.Println(message);
		fmt.Printf("Your new balance is ₹%d",balance);
	}

	
}

func withdraw(bal *int, withdrawAmt int) (string,error){
	if *bal < withdrawAmt {
		return "", fmt.Errorf("Your balance is ₹%d, which is insufficient to withdraw amount ₹%d.",*bal,withdrawAmt)
	}else{
		*bal = *bal - withdrawAmt
		return "Your transaction was successful.",nil
	}
}