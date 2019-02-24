package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r*c.r
}

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func distance(args... float64) float64 {
	total := float64(0)
	for _, v := range args {
		total += v
	}
	return total
}

type Person struct {
	Name string
}
func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

type Android struct {
	Person
	Model string
}

type Shape interface {
	area() float64
}

func main()  {

	c := Circle{x: 0, y: 0, r: 5}
	fmt.Println(c.x, c.y, c.r)
	fmt.Println(circleArea(&c))
	fmt.Println(c.area())

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())

	p := Person {Name:"Willy"}
	a := Android{Person: p, Model: "ABC100"}
	a.Person.Talk()


}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

