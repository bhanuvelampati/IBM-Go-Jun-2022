package main

import "fmt"

type Product struct {
	Id       int
	Name     string
	Cost     float32
	Units    int
	Category string
}

//implementing fmt.Stringer interface
func (p Product) String() string {
	return fmt.Sprintf("Id=%d, Name=%s, Cost=%v, Units=%d, Category=%s", p.Id, p.Name, p.Cost, p.Units, p.Category)
}

/*
	Write the apis for the following
		Sort => Sort the products collection by any attribute
			IMPORTANT : MUST Use sort.Sort() function
            use cases:
                1. sort the products collection by cost
                2. sort the products collection by name
                3. sort the products collection by units
                4. sort the products collection by cost in descending order
                5. sort the products collection by name in descending order
                6. sort the products collection by units in descending order
				etc

*/

type Products []Product

//implementing fmt.Stringer interface
func (products Products) String() string {
	result := ""
	for _, product := range products {
		result += fmt.Sprintf("%v\n", product)
	}
	return result
}

func (products Products) IndexOf(product Product) int {
	for idx, p := range products {
		if p == product {
			return idx
		}
	}
	return -1
}

func (products Products) Includes(product Product) bool {
	for _, p := range products {
		if p == product {
			return true
		}
	}
	return false
}

//Day-1
func (products Products) Filter(criteria func(Product) bool) Products {
	result := Products{}
	for _, product := range products {
		if criteria(product) {
			result = append(result, product)
		}
	}
	return result
}

func (products Products) All(criteria func(Product) bool) bool {
	for _, product := range products {
		if !criteria(product) {
			return false
		}
	}
	return true
}

func (products Products) Any(predicate func(Product) bool) bool {
	for _, product := range products {
		if predicate(product) {
			return true
		}
	}
	return false
}

type ProductPredicate func(product Product) bool

func Or(leftPredicate ProductPredicate, rightPredicate ProductPredicate) ProductPredicate {
	return func(product Product) bool {
		return leftPredicate(product) || rightPredicate(product)
	}
}

func And(leftPredicate ProductPredicate, rightPredicate ProductPredicate) ProductPredicate {
	return func(product Product) bool {
		return leftPredicate(product) && rightPredicate(product)
	}
}

func main() {
	products := Products{
		Product{105, "Pen", 5, 50, "Stationary"},
		Product{107, "Pencil", 2, 100, "Stationary"},
		Product{103, "Marker", 50, 20, "Utencil"},
		Product{102, "Stove", 5000, 5, "Utencil"},
		Product{101, "Kettle", 2500, 10, "Utencil"},
		Product{104, "Scribble Pad", 20, 20, "Stationary"},
		Product{109, "Golden Pen", 2000, 20, "Stationary"},
	}

	fmt.Println("Initial List")
	//fmt.Println(products.Format())
	fmt.Println(products)

	stove := Product{102, "Stove", 5000, 5, "Utencil"}
	fmt.Println("IndexOf stove = ", products.IndexOf(stove))
	fmt.Println("Does products include stove ?:", products.Includes(stove))

	fmt.Println()
	fmt.Println("Filter")
	fmt.Println("Costly Products [cost > 1000]")
	costlyProductPredicate := func(product Product) bool {
		return product.Cost > 1000
	}
	costlyProducts := products.Filter(costlyProductPredicate)
	//fmt.Println(costlyProducts.Format())
	fmt.Println(costlyProducts)

	fmt.Println("Stationary products [category = Stationary]")
	stationaryProductPredicate := func(product Product) bool {
		return product.Category == "Stationary"
	}
	stationaryProducts := products.Filter(stationaryProductPredicate)
	//fmt.Println(stationaryProducts.Format())
	fmt.Println(stationaryProducts)

	fmt.Println("Costly Stationary products")
	costlyStationaryProductPredicate := And(costlyProductPredicate, stationaryProductPredicate)
	costlyStationaryProducts := products.Filter(costlyStationaryProductPredicate)
	//fmt.Println(costlyStationaryProducts.Format())
	fmt.Println(costlyStationaryProducts)

	fmt.Println("All")
	fmt.Println("Are all products costly products ?:", products.All(costlyProductPredicate))

	fmt.Println("Any")
	fmt.Println("Are there any costly products ?:", products.Any(costlyProductPredicate))

	//Day-10
	utencilProducts := products.Filter(func(product Product) bool {
		return product.Category == "Utencil"
	})
	//fmt.Println(utencilProducts.Format())
	fmt.Println(utencilProducts)
}
