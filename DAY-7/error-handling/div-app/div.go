package main

import (
	"errors"
	"fmt"
)

func main(){
	
	var num1,num2 int
	fmt.Println("Enter 1st number")
	fmt.Scan(&num1)

	fmt.Println("Enter 2nd number")
	fmt.Scan(&num2)
	
	result,err := div(num1,num2)

	if(err != nil){
		fmt.Println("Error :",err)
	}else{
		fmt.Printf("result of %d / %d = %d",num1,num2,result)
	}
	
}

func div(a,b int) (int,error) {
	if b == 0{
		return 0, errors.New("can't divide by zero")
	}else{
		return a/b, nil
	}
}