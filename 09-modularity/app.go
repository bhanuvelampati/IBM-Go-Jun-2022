package main

import (
	"fmt"
	"modularity-demo/calculator"

	/* mathutils "modularity-demo/calculator/utils" */
	_ "modularity-demo/calculator/utils"

	"github.com/fatih/color"
)

func main() {
	//fmt.Println("modularity-demo executed")
	color.Red("modularity-demo executed")
	fmt.Println(calculator.Add(100, 200))
	fmt.Println(calculator.Subtract(100, 200))
	//fmt.Printf("operation count = %d\n", calculator.OpCount())
	color.Cyan("operation count = %d\n", calculator.OpCount())
	//fmt.Println("Is 21 even no ? :", mathutils.IsEven(21))
}
