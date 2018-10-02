package main

import "fmt"

func main() {

	// Various ways of declaring variables

	//var a int = 10
	//var b int = 20

	// var (
	// 	a int = 10
	// 	b int = 20
	// )

	//var a, b int = 10, 20

	// a := 10
	// b := 20

	a, b := 10, 20
	x, y := 10.4, 20.3333
	name := "John Doe"
	isBool := true

	fmt.Println("a,b :", a, b)
	fmt.Println("x,y :", x, y)
	fmt.Println("name :", name)
	fmt.Println("isBool :", isBool)

}
