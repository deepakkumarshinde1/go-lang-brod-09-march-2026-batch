package main

import (
	"errors"
	"fmt"
)

// error => interface inbuilt

func getIntegerValue(num *int) error {
	_,err := fmt.Scan(num)
	if(err != nil){
		return errors.New("Invalid number passed")
	}else{
		return nil
	}
}
func main(){
	var number int
	err:= getIntegerValue(&number)
	if err != nil{
		fmt.Println("error::",err)
	}else{
		fmt.Println("new number",number)
	}
	
}
	
