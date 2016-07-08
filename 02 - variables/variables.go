package main

import "fmt"

func main() {

	var a string = "initial" // Declare a variable
	var b, c int = 1, 2      // You can declare multiple variables
	var d = true             // Go will infer the type if initialized
	var e int                // Uninitialized variables are "zero-valued"
	f := "short"             // := is shorthand for declaration and initialization

	fmt.Println(a)
	fmt.Println(b, c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}
