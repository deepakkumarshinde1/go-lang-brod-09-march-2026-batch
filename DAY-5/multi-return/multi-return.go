package main

import "fmt"

func main(){
	sub,add := addAndSub(10,5)
	fmt.Println("Sub is", sub)
	fmt.Println("Add is", add)
}

// func Name(para) (returnType,returnType){}
func addAndSub(a,b int) (int,int) {
	sResult := a-b
	aResult := a+b
	return sResult,aResult
}