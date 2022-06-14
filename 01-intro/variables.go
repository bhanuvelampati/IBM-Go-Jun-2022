package main

import "fmt"

func main() {
	/*
		var msg string
		msg = "Hello World"
	*/
	/*
		var msg string = "Hello World!"
	*/

	/*
		//type inference
		var msg = "Hello World!"
	*/
	msg := "Hello World!" // := syntax is applicable ONLY in function scope (NOT in package level)
	fmt.Println(msg)

	//multiple variable declarations
	/*
		var x int
		var y int
		var result int
		var str string
		x = 100
		y = 200
		str = "Sum of 100 and 200 is"
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var x int = 100
		var y int = 200
		var result int = x + y
		var str string = "Sum of 100 and 200 is"
		fmt.Println(str, result)
	*/

	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		str = "Sum of 100 and 200 is"
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		str = "Sum of 100 and 200 is"
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y   int = 100, 200
			result int
			str    string = "Sum of 100 and 200 is"
		)
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		var (
			x, y, result = 100, 200, 0
			str          = "Sum of 100 and 200 is"
		)
		result = x + y
		fmt.Println(str, result)
	*/

	/*
		x, y, result := 100, 200, 0
		str := "Sum of 100 and 200 is"

		result = x + y
		fmt.Println(str, result)
	*/

	x, y, str := 100, 200, "Sum of 100 and 200 is"
	result := x + y
	fmt.Println(str, result)
}
