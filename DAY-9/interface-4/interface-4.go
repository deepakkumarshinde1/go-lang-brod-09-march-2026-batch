package main

import "fmt"

func print(varOne interface{}) {
	value, ok := varOne.(int)
	if ok == false {
		fmt.Println("Value must be integer")
	} else {
		fmt.Println(value)
	}
}
func main() {
	var number interface{}
	// value , type
	number = 10
	print(number)

}
