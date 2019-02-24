package main

import "fmt"

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}

func main() {

	defer second()
	first()

	xs := []float64{98,93,77,82,83}

	fmt.Println(average(xs))

	x, y := f()

	fmt.Println(x, y, " Result ")

	ys := []int{1,2,3}
	fmt.Println(add(ys...))

	a := 0
	increment := func() int {
		a++
		return a
	}
	fmt.Println(increment())
	fmt.Println(increment())

	fmt.Println(factorial(5))

}

func average(xs []float64) float64 {

	total := 0.0
	for _, v := range xs {
		total += v
	}
	return total / float64(len(xs))

}

func add(args ...int) int {
	total := 0
	for _, v := range args {
		total += v
	}
	return total
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func f() (int, int) {
	return 5, 6
}