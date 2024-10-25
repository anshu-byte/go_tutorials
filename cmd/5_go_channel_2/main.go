package main

import "fmt"

func main() {
    channel := make(chan string)
	go process(channel)
    fmt.Println(<-channel)
    fmt.Println("Finished")
}

func process(channel chan string) {
	channel <- "forty two"
} 