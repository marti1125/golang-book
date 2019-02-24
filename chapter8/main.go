package main

import "fmt"

func zero(x int) {
	x = 0
}

func zeroPointer(xPtr *int) {
	*xPtr = 0
}

func one(xPtr *int) {
	*xPtr = 1
}

func square(a *float64) {
	*a = *a * *a
}

func main()  {

	x := 5
	zero(x)
	fmt.Println(x) // x is still 5

	zeroPointer(&x)
	fmt.Println(x) // x is 0

	xPtr := new(int)
	one(xPtr)
	fmt.Println(*xPtr) // x is 1

	a := 1.5
	square(&a)

	fmt.Println(a)

}
