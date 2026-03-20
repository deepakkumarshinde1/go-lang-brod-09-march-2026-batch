package main

import (
	"fmt"
	"time"
)

func printName(name string) {
	fmt.Println(name)
}
func main() {
	timer := time.Now()

	go printName("Deepakkumar")
	go printName("Suraj")
	go printName("Archana")
	go printName("Smita")

	time.Sleep(time.Microsecond * 50)
	fmt.Println(time.Since(timer))
}
