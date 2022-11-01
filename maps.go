package main

import (
	"fmt"
)

func main() {
	var x = make(map[string]string)
	x["nam1"] = "abid"
	x["nam2"] = "shah"
	x["nam3"] = "mhmd"

	fmt.Println(x)

	delete(x, "nam3")

	fmt.Println(x)
	a := map[string]int{"one": 1, "two": 2, "three": 3, "four": 4}

	var b = []string{} // defining the order
	b = append(b, "one", "two", "three", "four")

	for k, v := range a { // loop with no order
		fmt.Printf("%v : %v, ", k, v)
	}

	fmt.Println()

	for _, element := range b { // loop with the defined order
		fmt.Printf("%v : %v, ", element, a[element])
	}
}
