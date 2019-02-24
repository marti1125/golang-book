package main

import "fmt"

var z string = "Hello World"

var (
	a = 5
	b = 10
	c = 15
)

func main()  {

	var x string = "Hello World"
	fmt.Println(x)

	var second string
	second = "second"
	fmt.Println(second)
	second = "first"
	fmt.Println(second)

	var a string = "hello"
	var b string = "world"

	fmt.Println(a == b)

	// shorter statement
	v := "Hello World"
	fmt.Println(v)

	y := 5
	fmt.Println(y)

	dogsName := "Max"
	fmt.Println("My dog's name is", dogsName)

	f()

	const c string = "Hello World"
	fmt.Println(c)

}

func f() {
	fmt.Println(z)
}