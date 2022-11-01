package main

import (
	"fmt"
)

func main() {
	var (
		a int
		b = 3
	)
	c, d := 7, "World!"

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println("shahensahh.....")

	/*Floating Numbers*/

	var i = 3.141

	fmt.Printf("%e\n", i)
	fmt.Printf("%f\n", i)
	fmt.Printf("%.2f\n", i)
	fmt.Printf("%6.2f\n", i)
	fmt.Printf("%g\n", i)

	/*Boolean*/
	var m = true
	var n = false

	fmt.Printf("%t\n", m)
	fmt.Printf("%t\n", n)

	/*String values*/

	var txt = "Hello"

	fmt.Printf("%s\n", txt)
	fmt.Printf("%q\n", txt)
	fmt.Printf("%8s\n", txt)
	fmt.Printf("%-8s\n", txt)
	fmt.Printf("%x\n", txt)
	fmt.Printf("% x\n", txt)

	/*Integer representions*/

	var k = 15

	fmt.Printf("%b\n", k)
	fmt.Printf("%d\n", k)
	fmt.Printf("%+d\n", k)
	fmt.Printf("%o\n", k)
	fmt.Printf("%O\n", k)
	fmt.Printf("%x\n", k)
	fmt.Printf("%X\n", k)
	fmt.Printf("%#x\n", k)
	fmt.Printf("%4d\n", k)
	fmt.Printf("%-4d\n", k)
	fmt.Printf("%04d\n", k)

}
