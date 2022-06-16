package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

type PerishableProduct struct {
	Product
	/* 	Dummy */
	Expiry string
}

type Dummy struct {
	Cost float32
}

func NewPerishableProduct(id int, name string, cost float32, units int, category string, expiry string) PerishableProduct {
	return PerishableProduct{
		Product: Product{
			Id:       id,
			Name:     name,
			Cost:     cost,
			Units:    units,
			Category: category,
		},
		Expiry: expiry,
	}
}

func main() {
	/* grapes := PerishableProduct{
		Product{101, "Grapes", 50, 10, "Food"},
		"5 Days",
	} */

	/*
		grapes := PerishableProduct{
			Product: Product{101, "Grapes", 50, 10, "Food"},
			Expiry:  "5 Days",
		}
	*/
	grapes := NewPerishableProduct(101, "Grapes", 50, 10, "Food", "5 Days")
	fmt.Println(grapes)
	fmt.Println("grapes Expiry = ", grapes.Expiry)
	//fmt.Println("grapes Cost =", grapes.Product.Cost)
	fmt.Println("grapes Cost =", grapes.Cost)
}
