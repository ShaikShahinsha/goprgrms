package main

import (
	"fmt"
)

func main() {
	var arr1 = [...]string{"shah", "abid"}
	arr2 := [...]string{"abid", "ameen", "chunu", "sofi", "asadh", "mohmd"}
	fmt.Println(arr1)
	fmt.Println(arr2)

	sval := []int{}
	slic3 := arr2[1:4]
	fmt.Println(len(sval))
	fmt.Println(len(arr1))
	fmt.Println(cap(arr1))
	fmt.Println(slic3)
	fmt.Println(len(slic3))

	myslice1 := make([]int, 5, 10)
	fmt.Printf("myslice1 = %v\n", myslice1)

	//append elements of a slice

	myslice := []int{1, 2, 3, 4, 5, 6}
	fmt.Printf("myslice = %v\n", myslice)
	myslice = append(myslice, 50, 88, 9965)
	fmt.Printf("myslice = %v\n", myslice)

	//append one slice to another....

	myslic3 := append(myslice, myslice1...)
	fmt.Printf("myslice3 = %v\n", myslic3)

	//copy() for memory optimization..

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	fmt.Printf("numbers = %v\n", numbers)
	// Create copy with only needed numbers
	neededNumbers := numbers[:len(numbers)-10]
	numbersCopy := make([]int, len(neededNumbers))
	copy(numbersCopy, neededNumbers)

	fmt.Printf("numbersCopy = %v\n", numbersCopy)
}
