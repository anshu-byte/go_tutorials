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
	for i:= range c{
		fmt.Println(i)
	}
}

func process(c chan int){
	defer close(c)
	// c <- 123
	for i:=0; i<5; i++{
		c <- i
	}
}