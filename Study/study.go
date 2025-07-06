package main

import "fmt"

func loop() {
	for i := 1; i <= 10; i++ {
		fmt.Print(i, " ")
	}
}
func main() {
	loop()
	var name = "Hello"
	secName := "World"

	fmt.Println("\n", name, secName)

	var (
		age = 20
		sex = "male"
	)

	fmt.Println(age, sex)

	var a, b, c int = 1, 2, 3

	fmt.Println(a, b, c)

	var d = 1
	var e = 2
	var f = 3

	fmt.Println(d, e, f)
}
