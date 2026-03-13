package main

import "fmt"

func main(){
		fmt.Println("Working ...")
		response := make([]int,10)
		responseApiCount := 10
		var responseAvg int;
		var slowAPis int

		for i := range responseApiCount{
			fmt.Scan(&response[i])
		}

		for i := range response{
			responseAvg += response[i]
			if response[i] > 500{
				slowAPis++
			}
		}
		
		responseAvg = responseAvg/len(response);

		fmt.Printf("Average response time: %dms \n",responseAvg)
		fmt.Println("Slow APIs: ",slowAPis)
		


}