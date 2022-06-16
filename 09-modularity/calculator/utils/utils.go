package utils

import "fmt"

var ProductsRank map[string]int

func init() {
	fmt.Println("utils package initialized")
	ProductsRank = make(map[string]int)
}

func IsEven(no int) bool {
	return no%2 == 0
}
