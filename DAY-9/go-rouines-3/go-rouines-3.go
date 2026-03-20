package main

import (
	"fmt"
	"sync"
)

// waitGroup
func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(1)

	go func() {
		fmt.Println("Hello world")
		waitGroup.Done()
	}()

	fmt.Println("Last Line")
	waitGroup.Wait()

}
