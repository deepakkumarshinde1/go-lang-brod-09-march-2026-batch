package main

import "fmt"


func main(){
	sub,add := addAndSub(10,5)
	fmt.Println("Sub is", sub)
	fmt.Println("Add is", add)

	text := print("Deepakkumar")
	fmt.Println(text)
}

// func Name(para) (varName returnType,rName returnType){}
func addAndSub(a,b int) (sResult int, aResult int) {
	sResult = a-b
	aResult = a+b
	return
}

func print(text string) (newText string){
		newText = "Hello " +text
		return
}