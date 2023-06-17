package main

import (
	"fmt"
	//"rsc.io/quote"
)

func main() {
	//fmt.Println("Hello, World!")
	// var name string = "Lakshan Rukantha"
	// var age int = 22

	// var name = "Lakshan"

	// fmt.Println(name)

	// name := "Lakshan Rukantha"
	// age := 22
	// weight := 55.25

	// fmt.Printf("I am %v and I am %v years old.\n", name, age)
	// fmt.Printf("Weight: %g Kg", weight)

	//fmt.Println(quote.Go())

	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println("Counting...", i+1)
	}

	//numArray := [5]int{1,2,3,4,5}

	//fruits := [5]string{"Apple", "Orange", "Mango", "Banana", "Grapes"}

	// fmt.Println("I like", fruits[3])

	// for i := 0; i < 5; i++ {
	// 	fmt.Println(numArray)
	// }

	// // Print 2d array

	// limit := 3

	// for i := 0; i < limit; i++ {
	// 	for j := 1; j <= limit; j++ {
	// 		fmt.Print(i*3 + j)
	// 	}
	// 	fmt.Println("")
	// }

}