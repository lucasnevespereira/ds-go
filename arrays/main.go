package main

import "fmt"

func main() {
	// fixed len
	var myArray [8]int
	fmt.Println(myArray)

	// non-fixed len
	var mySlice []int
	mySlice = append(mySlice, 1, 2, 3, 4)
	fmt.Println(mySlice)
}
