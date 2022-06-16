package calculator

import "fmt"

func init() {
	fmt.Println("calculator intialized - [subtract.go]")
}

func Subtract(x, y int) int {
	opCount++
	return x - y
}
