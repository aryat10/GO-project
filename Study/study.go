package main

import "fmt"

// func loop() {
// 	for i := 1; i <= 10; i++ {
// 		fmt.Print(i, " ")
// 	}
// }
func main() {
	// loop()
	// var name = "Hello"
	// secName := "World"

	// fmt.Println("\n", name, secName)

	// var (
	// 	age = 20
	// 	sex = "male"
	// )

	// fmt.Println(age, sex)

	// var a, b, c int = 1, 2, 3

	// fmt.Println(a, b, c)

	// var d = 1
	// var e = 2
	// var f = 3

	// fmt.Println(d, e, f)

	// we can declare a variable inside if construct
	// if age := 20; age > 18 {
	// 	fmt.Println("You are an adult")
	// } else {
	// 	fmt.Println("You are a minor")
	// }

	// we can declare a variable inside switch construct
	switch age := 20; age {
	case 18:
		fmt.Println("You are an adult")
	case 19:
		fmt.Println("You are a minor")
	case 20:
		fmt.Println("You are an adult")
	default:
		fmt.Println("You are a minor")
	}
}
