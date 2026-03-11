package main

import "fmt"

func user() (string, string) {
	name := "deepakkumar"
	location := "nashik"
	// fmt.Println(2 + 2, name)
	return name, location
}

// func add(a int, b int) int {
//	return a + b
// }

func add(a, b int) int {
	return a + b
}

func div(a, b int) (int, error) {

	if b == 0 {
		return 0, fmt.Errorf("Can't div by zero")
	}
	result := a / b
	return result, nil
}

var multi = func(a, b int) int {
	return a * b
}

func increment(start int) func() {
	counter := start - 1
	return func() {
		counter++
		fmt.Println("counter is ", counter)
	}
}

func mainOne() {

	myInc := increment(1)

	myInc()
	myInc()
	myInc()
	myInc()
	myInc()

	result := add(10, 20)
	fmt.Println(result)
	divResult, divError := div(10, 10)

	multiResult := multi(10, 20)

	fmt.Println(multiResult)
	if divError == nil {
		fmt.Println(divResult)
	} else {
		fmt.Println(divError)
	}

	// name := "Deepakkumar"

	// var number int = 10 // 32 bit or 64 bit
	// var number1 int8 = 11 // -128 to 127
	// var number2 int16 = 11 // -32,768  to 32,767
	// var number3 int32 = 11 // -2,147,483,648 to 2,147,483,647
	// var number4 int64 = 11 // -9,223,372,036,8544,08 to 9,223,372,036,8544,09

	// var f1 float32 = 10.10 // ~6 decimal digit
	// var f1 float64 = 10.10 // ~15 decimal digit

	// var name string = "nashik"
	// var isPresent = true // true or false

	// var n1 uint = 200 // 32 bit or 64 bit
	// var n2 uint8 = 200 // 0 to 255
	// var n5 byte = 200 // 0 to 255 alias for uint8
	// var n3 uint16 = 200 // 0 to 65,535
	// var n4 uint32 = 200 // 0 to 4,294,967,295
	// var n5 uint64 = 20 // 0 to 18,446,744,073,709,615

	// fmt.Println(number)
	// fmt.Println("Hello", name)

	// IIF Immediately Invoked Function
	// (func(a, b int) {
	//fmt.Println(a + b)
	//})(10, 400)

}

func mainTwo() {
	var numberOne, numberTwo int
	fmt.Println("Enter a number: ")
	fmt.Scanf("%d %d", &numberOne, &numberTwo)
	fmt.Printf("Add is  of %d + %d is %d", numberOne, numberTwo, numberOne+numberTwo)
}

// a,b check for b === 0 return error or return  result

// use  a function like div

func main() {

}