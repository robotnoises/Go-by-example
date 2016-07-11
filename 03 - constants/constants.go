package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {

	const n = 500000000
	const d = 3e20 / n

	fmt.Println(s)
	fmt.Println(d)           // constants perform arithmetic with arbitrary precision
	fmt.Println(int64(d))    // explicit cast
	fmt.Println(math.Sin(n)) // implicit type via func (also assignment)
}
