package main

import "fmt"

// How we can handle a slice or array with pointer.
func main(){
	list := [4]int{100,200,300,400}
	fmt.Println("outer list before",list)
	updateList(&list)
	fmt.Println("outer list after",list)
	
}
// I the slices are auto references. So you don't need to use pointer over here.
// For arrays, you need to use a address operator and Pointer type but No de-rereferencing operator.
func updateList(l *[4]int){
	l[0] = 1234
	fmt.Println("inner list:",*l)
}