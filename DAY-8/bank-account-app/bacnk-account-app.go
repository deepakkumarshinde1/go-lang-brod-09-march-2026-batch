package main

import (
	"errors"
	"fmt"
)

type BankAccount struct{
	ownerName string
	balance int
}

func (b *BankAccount) createAccount(name string, initialBalance int){
	b.ownerName = name
	b.balance = initialBalance
	fmt.Println("\nAccount Created successfully !!!")
}

func (b BankAccount) displayAccountDetails(){
	fmt.Println("\n---- Here are your bank details ----")
	fmt.Println("\nAccount Holder :",b.ownerName)
	fmt.Println("\nAccount Balance :",b.balance)
}

func (b *BankAccount) withdrawAmount(amount int) error{
	if b.balance < amount{
		return errors.New("Insufficient amount.")
	}else{
		b.balance -= amount
		fmt.Printf("Your amount ₹%d has been withdraw successfully and your current balance is ₹%d.",amount,b.balance)
		return nil
	}
}

func (b BankAccount) checkBalance(){
		fmt.Println("\nYour balance is:",b.balance)
}
func main(){
	ba := BankAccount{} 

	ba.createAccount("Deepakkumar",1000)
	ba.displayAccountDetails()
	err := ba.withdrawAmount(600)

	if(err != nil){
		fmt.Println(err)
	}

	ba.checkBalance()

	// fmt.Printf("\n%+v \n", ba)


}