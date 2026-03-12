package main

import "fmt"

// slice
var taskList = []string{}

func addNewTask(){
	fmt.Println("\n Enter a task")
	var newTask string
	fmt.Scan(&newTask)
	taskList = append(taskList,newTask)
	fmt.Println("\n task saved !!!")
	menu()
	
}

func printTask(){
	fmt.Println("\n Here is your task")
	for i,value := range taskList{
		fmt.Println(i+1,". ", value)
	}
	menu()
}

func menu(){
	fmt.Println("\n1. Create Task")
	fmt.Println("2. My Task")
	fmt.Println("3. Exit")

	var key int

	fmt.Scan(&key)

	switch key{
	case 1:
		addNewTask()
	case 2:
		printTask()
	case 3:
		fmt.Println("Thank you...")
	default:
		fmt.Println("\nError:: Invalid option, try again.")
		menu()
	}

}


func main(){

	fmt.Println("Welcome to task app")
	menu()
	
}