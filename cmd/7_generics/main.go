package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type contactInfo struct {
	Name  string
	Email string
}

type purchaseInfo struct {
	Name   string
	Price  float32
	Amount int
}

func main() {
	contacts := loadJSON[contactInfo]("contacts.json")
	fmt.Printf("%v\n", contacts)
	
	purchases := loadJSON[purchaseInfo]("purchases.json")
	fmt.Printf("%v\n", purchases)
}

func loadJSON[T any](filePath string) []T {
	data, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Failed to read file: %v\n", err)
		return nil
	}

	var loaded []T
	err = json.Unmarshal(data, &loaded)
	if err != nil {
		fmt.Printf("Failed to parse JSON: %v\n", err)
		return nil
	}
	return loaded
}
