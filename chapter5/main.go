package main

import "fmt"

func main()  {

	i := 1

	for i <= 10 {
		if i % 2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
		i += 1
	}

	for a := 1; a <= 10; a++ {
		if a == 0 {
			fmt.Println("Zero")
		} else if a == 1 {
			fmt.Println("One")
		} else if a == 2 {
			fmt.Println("Two")
		} else if a == 3 {
			fmt.Println("Three")
		} else if a == 4 {
			fmt.Println("Four")
		} else if a == 5 {
			fmt.Println("Five")
		}
	}

	b := 2

	switch b {
		case 0: fmt.Println("Zero")
		case 1: fmt.Println("One")
		case 2: fmt.Println("Two")
		case 3: fmt.Println("Three")
		case 4: fmt.Println("Four")
		case 5: fmt.Println("Five")
		default: fmt.Println("Unknown Number")
	}

}