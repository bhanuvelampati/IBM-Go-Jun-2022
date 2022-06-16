package calculator

import "fmt"

func init() {
	fmt.Println("calculator intialized - [add.go]")
}
func Add(x, y int) int {
	opCount++
	return x + y
}
