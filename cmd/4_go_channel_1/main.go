package main

import "fmt"

// Go channels
// Hold Data, Thread Safe and Listen for Data

func main(){
	var c = make(chan int)
	// c <- 1
	// var i = <- c // reading should be done after the go routine i.e with process function
	// fmt.Println(i)
	go process(c)

	// if you range over channel then you have to specify defer close(c) in the process function
	// for i:= range c{
	// 	fmt.Println(i)
	// }

	// better way to avoid hanging out the channel
	for i:=0; i<5; i++{
		x := <- c
		fmt.Println(x)
	}
}

func process(c chan int){
	// defer close(c)
	// c <- 123
	for i:=0; i<5; i++{
		c <- i
	}
}