package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bs, err := ioutil.ReadFile("./chapter13/test.txt")
	if err != nil {
		fmt.Println("Error....")
		return
	}
	str := string(bs)
	fmt.Println(str)
}