package main

import (
	"fmt"
	"strings"
)


type gasEngine struct{
	mpg uint8
	gallons uint8 // we can even use function as a type like calculateGas uint8
}

type electricEngine struct{
	mpkwh uint8
	kwh uint8
}

func (e gasEngine) milesLeft() uint8{
	return e.gallons * e.mpg
}
	
func (e electricEngine) milesLeft() uint8{
	return e.kwh * e.mpkwh
}

type engine interface{
	milesLeft() uint8
}

func canMakeIt(e engine, miles uint8){
	if miles <= e.milesLeft(){
		fmt.Println("We can make it")
	}else{
		fmt.Println("We can't make it")
	}
}
func main(){
	 var myString string = "resume"
	 var indexed = myString[0]
	 fmt.Println(indexed)

	 for i, v := range myString{
		 fmt.Println(i, v)
	 }

	// strings are immutable in go
	// myString[0] = 'a' // error

	var strSlice = []string{"one", "two", "three", "four", "five"}
	var catStr = ""
	for i := range strSlice{
		catStr += strSlice[i] // inefficient for large strings, cause it creates new string
	}
	fmt.Println(catStr)

	// efficient for large strings
	var strBuilder strings.Builder
	for i := range strSlice{
		strBuilder.WriteString(strSlice[i])
	}
	fmt.Println(strBuilder.String())

	var myEngine gasEngine = gasEngine{25, 5}
	canMakeIt(myEngine, 20)

	var myEngine2 electricEngine = electricEngine{25, 5}
	canMakeIt(myEngine2, 20)
}