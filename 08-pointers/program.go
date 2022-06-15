package main

import "fmt"

func main() {
	var no int = 10
	var noPtr *int = &no
	fmt.Println(noPtr, no)

	//dereferencing - access the underlying value at an address
	var valueAtPtr = *noPtr
	fmt.Println(valueAtPtr)

	fmt.Println("Before incrementing, no = ", no)
	increment(&no)
	fmt.Println("After incrementing, no = ", no)

	x, y := 100, 200
	fmt.Printf("Before swapping, x = %d and y = %d\n", x, y)
	swap( /*  */ )
	fmt.Printf("After swapping, x = %d and y = %d\n", x, y)

}

func increment(x *int) {
	*x += 1
}

func swap( /*  */ ) {

}
