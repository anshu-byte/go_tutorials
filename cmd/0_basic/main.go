package main

import (
	"errors"
	"fmt"
	"time"
	"unicode/utf8"
)

func main(){
	fmt.Println("Hello World!")
	var intNum int = 32767 // 32767 -> highest value 'int' can hold
	fmt.Println(intNum)

	var floatNum float32 = 123.45
	fmt.Println(floatNum)
	
	// You can't do mathematical operations on different type

	var floatNum32 float32 = 10.1
	var intNum32 int32 = 2
	var result float32 = floatNum32 + float32(intNum32)
	fmt.Println(result)

	var myString string = "Hello" + " " + "World"
	fmt.Println(myString)
	
	fmt.Println(len("test")) // number of bytes 
	fmt.Println(utf8.RuneCountInString("test")) // length of string

	// default value of all type of 'int' is 0
	// default value of string is ''
	// default value of boolean is false

	var myBoolean bool = false
	fmt.Println(myBoolean)

	var myVar = "text" // type is inferred here
	fmt.Println(myVar)

	myVar2 := "text"
	fmt.Println(myVar2)

	var1, var2 := 1, 2
	fmt.Println(var1, var2)

	const myConst string = "const value"
	fmt.Println(myConst)

	var printValue string = "Hello World"
	printMe(printValue)

	var numerator int = 14
	var denominator int = 3
	var result2, remainder, err = intDivision(numerator, denominator)
	if err != nil{
		fmt.Println(err.Error())
	}else if remainder == 0{
		fmt.Printf(`The result of the integer division is %v`, result2)
	}else{
		fmt.Printf(`The result of the integer division is %v with remainder %v`, result2, remainder)
	}
	fmt.Println("")
	fmt.Println("Same thing by using switch")
	switch{
		case err != nil:
			fmt.Println(err.Error())
		case remainder == 0:
			fmt.Printf(`The result of the integer division is %v`, result2)
		default:
			fmt.Printf(`The result of the integer division is %v with remainder %v`, result2, remainder)
	}

	var intArr [3]int32 // int32 -> 4 bytes
	intArr[1] = 123
	fmt.Println(intArr[0])
	fmt.Println(intArr[1:3])

	fmt.Println(&intArr[0])
	fmt.Println(&intArr[1])
	fmt.Println(&intArr[2])

	// shortcut 1
	var intArr2 [3]int32 = [3]int32{1, 2, 3}
	fmt.Println(intArr2)

	// shortcut 2
	intArr3 := [3]int32{1, 2, 3}
	fmt.Println(intArr3)

	// shortcut 3
	intArr4 := [...]int32{1, 2, 3, 4}
	fmt.Println(intArr4)


	var intSlice []int32 = []int32{4, 5, 6}
	fmt.Println(intSlice)
	fmt.Printf("The length is %v and capacity is %v\n", len(intSlice), cap(intSlice))
	intSlice = append(intSlice, 7)
	fmt.Printf("The length is %v and capacity is %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)

	// How we got 4 length intSlice arr?, we had define intSlice arr size to be 3
	// New intSlice arr will be created at different address

	var intSlice2 []int32 = []int32{8, 9, 10}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)
	fmt.Printf("The length is %v and capacity is %v\n", len(intSlice), cap(intSlice))

	// capacity becomes double when we append more elements than current capacity
	
	var intSlice3 []int32 = make([]int32, 3, 8)
	fmt.Printf("The length is %v and capacity is %v\n", len(intSlice3), cap(intSlice3))

	var myMap map[string]uint8 = make(map[string]uint8)
	fmt.Println(myMap) 

	var myMap2 = map[string]uint8 {"Adam": 23, "Sarah": 45, "Richard": 67}
	fmt.Println(myMap2)
	fmt.Println(myMap2["Adam"])
	fmt.Println(myMap2["Sarah"])
	fmt.Println(myMap2["Anshu"]) // Map will return 0 if key doesn't exist

	var value, ok = myMap2["Anshu"]
	if !ok{
		fmt.Println("Name doesn't exist")
	}else{
		fmt.Println(value)
	}

	delete(myMap2, "Sarah")
	fmt.Println(myMap2)

	for name, age:= range myMap2{
		fmt.Printf("Name: %v, Age: %v \n", name, age)
	}

	for i:=0; i<3; i++ {
		fmt.Println(i)
	}

	var n int = 1000000
	var testSlice = []int{}
	var testSlice2 = make([]int, 0, n)
	fmt.Printf("Total time without preallocation: %v\n", timeLoop(testSlice, n))
	fmt.Printf("Total time with preallocation: %v\n", timeLoop(testSlice2, n))

	// best way for slice
	testSlice3 := make([]int, 0, n)
	for i:=0; i<5; i++{
		testSlice3 = append(testSlice3, i)
	}
	fmt.Println(testSlice3)

}

func timeLoop(slice []int, n int) time.Duration{
	var t0 = time.Now()
	for len(slice) < n{
		slice = append(slice, 1)
	}
	return time.Since(t0)
}

func printMe(printValue string){
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}