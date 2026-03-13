package main

import "fmt"

func main(){
	var productCount int
	var totalPrice int
	var discountPrice float64
	fmt.Println("Enter Product Count ")
	fmt.Scan(&productCount)

	products := make([]int,productCount)
	
	fmt.Printf("Enter %d product amount \n", productCount)

	for i := range products{
		fmt.Print("Enter Product ",i+1," Amount: ")
		fmt.Scan(&products[i])
		fmt.Println("")
	}

		for _,value := range products{
			totalPrice += value
		}

		if totalPrice > 5000{
			discountPrice = float64(totalPrice) * 10/100
		}

		fmt.Println("Total: ₹", totalPrice)
		fmt.Println("Discount: ₹",discountPrice)
		fmt.Println("Final Price: ₹" , totalPrice - int(discountPrice))
		


	

}