package main

import "fmt"

func main()  {

	fmt.Println("Fahrenheit to degrees Celsius")
	fmt.Print("Enter a number: ")
	var input float64
	fmt.Scanf("%f", &input)

	fmt.Println(FahrenheitIntoCelsius(input),  "C")

}

func FahrenheitIntoCelsius(f float64) float64 {
	return (f - 32) * 5/9
}