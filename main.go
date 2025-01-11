package main

import (
	"fmt"
)

func add(num1 int, num2 int) {
	var sum = num1 + num2

	fmt.Println("This is total", sum)

}

func main() {
	// fmt.Println("hello world")

	// a := 12
	// fmt.Println(a)

	// If else condition use
	/* var a = 18

	if a > 18 {
		fmt.Println("You are eligiable for married")
	} else if a < 18 {
		fmt.Println("You are not eligable not married")
	} else {
		fmt.Println("Now You are teenager OKH")
	}

	*/

	// var name = 243

	// if name >= 32 {
	// 	fmt.Println("This is biggest number")

	// } else if name <= 34 {
	// 	fmt.Println("This is lowest number", +name)
	// } else {
	// 	fmt.Println("Not this exucuted")
	// }
	//

	// Function learning

	const a = 19
	const b = 13
	add(a, b)

}
