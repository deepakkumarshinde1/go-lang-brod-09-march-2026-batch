package main

import (
	"fmt"
	"sync"
	"time"
)

func uploadFiles(fileName string, wg *sync.WaitGroup) bool {
	defer wg.Done()
	fmt.Printf("\nFile %v is uploading...\n", fileName)
	time.Sleep(2 * time.Second)
	fmt.Printf("File %v is uploaded\n", fileName)

	return true
}
func main() {
	files := []string{"file-1.png", "file-2.pdf", "file-3.jpeg"}
	var wg sync.WaitGroup

	for _, value := range files {
		wg.Add(1)
		go uploadFiles(value, &wg)
	}

	wg.Wait()

}
