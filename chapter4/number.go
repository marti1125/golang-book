package main

import "fmt"

func main()  {

	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	output := input * 2

	x := 5; x += 1

	fmt.Println(output)
	fmt.Println(x)

}
