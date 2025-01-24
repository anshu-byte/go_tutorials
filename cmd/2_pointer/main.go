package main

import "fmt"
func main(){
	var p *int32 // doesn't have an address assigned to it yet
	var i int32
	fmt.Println(p)

	// pointer is a variable that holds the address of another variable
	// assign memory pointer to p1
	var p1 *int32 = new(int32)
	fmt.Println(p1)
	fmt.Println(*p1)
	*p1 = 10
	fmt.Println(*p1)
	p1 = &i // assigns address of i to p1
	fmt.Println(*p1)
	i = 20
	fmt.Println(*p1)

	var slice = []int{1, 2, 3}
	var sliceCopy = slice
	sliceCopy[0] = 10
	fmt.Println(slice)
	fmt.Println(sliceCopy)
	// sliceCopy points to the same memory address as slice
	pointerUseCase()
}

func pointerUseCase(){
	var thing1 = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("\nThe memory location of the things1 array is %p\n", &thing1)
	var result [5]int = square(&thing1)
	fmt.Printf("The result is %v\n", result)
	fmt.Printf("The value of thing1 is %v\n", thing1)
}

func square(thing2 *[5]int) [5]int{
	fmt.Printf("The memory location of the thing2 array is %p\n", &thing2)
	for i := range thing2{
		thing2[i] = thing2[i] * thing2[i]
	}
	return *thing2
}