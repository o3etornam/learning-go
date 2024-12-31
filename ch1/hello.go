package main

import (
	"fmt"
	"reflect"
)

func main() {

	quantity := 4
	length, width := 1.2, 2.4
	customerName := "Damon Cole"

	fmt.Println(customerName)
	fmt.Println("has ordered", quantity, "sheets")
	fmt.Println("each with an area of", length*width, "square meters")

	fmt.Println(reflect.TypeOf(false))

}
