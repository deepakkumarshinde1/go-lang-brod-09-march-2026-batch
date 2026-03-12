package main

import "fmt"

func sliceOne_(){
	// var number []int
	// var number1 = []int{}

	// numbers := []int{10,20,30,40,50,60}

	numbers := []int{10,20}
	// add record
	// numbers = append(numbers,100)
	number := append(numbers,100,200,300,400,500,600)

	fmt.Println(numbers)
	for i,value := range number{
		fmt.Println(i,value)
	}


}

func sliceTwo_(){
	// numberOne := []int{10,20,30,40}
	// numberTwo := []int{}
	numberTwo := make([]int,0)
	
	numberTwo = append(numberTwo,100)

	// fmt.Println(len(numberOne))
	fmt.Println(numberTwo)
	// fmt.Println(cap(numberOne))
	// fmt.Println(len(numberTwo),cap(numberTwo))




}

func sliceThree_(){
	students := []string{"a","b","c","d"} 
	// copy
	var studList  = make([]string,len(students))
	
	copy(studList,students)
	students[0] = "Deepakkumar"
	fmt.Println(students)
	fmt.Println(studList)

	sliceOne := []int{10,30}
	sliceTwo := []int{10,20}
	
	// fmt.Println(arrayOne == arrayTwo)
	// camp

	isEquals := true 
	for i := range sliceOne{
		if len(sliceOne) != len(sliceTwo){
			isEquals = false
			break
		}
		if sliceOne[i] != sliceTwo[i]{
			isEquals = false
			break
		}
	}

	if isEquals {
		fmt.Println("Slices are equal")
	}else{
		fmt.Println("Slices are no equal")
	}

}

func sliceFour(){
	// slice opt (:)
	numbers := []int{10,20,30,40,50,60,70,80,90,100}

	// slice[start:end] => create slice form start index to end index
	// slice[start:] => create slice from given index to end
	// slice[:end] => create from start to index
	// slice[:] => create with all rec

	// slice := numbers[:]
	// slice := numbers[1:8]
	// slice := numbers[:5]
	slice := numbers[5:]
	// slice = append(slice,100)
	fmt.Println(slice)

}

func sliceFive(){
	studentsBatchOne := []string{"soham","amit","neha"}
	studentBatchTwo :=  []string{"deepakkumar","suraj","rohan"}

	studentBatch := append(studentsBatchOne,studentBatchTwo...)

	fmt.Println(studentBatch)

}

func main(){
		numbers := []int{10,20,30,40,50,60,70,80,90,100}
		// remove index = 0
		// numbers = numbers[1:]
		// numbers = numbers[: len(numbers)-1]
		index := 2
		numbers = append(numbers[:index], numbers[index+1:]...)
		fmt.Println(numbers)
}