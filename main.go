package main

import (
	"Algorithm/simhash"
	"fmt"
)

func main() {
	a := "abcdefg"
	b  := simhash.ComputeSimHashForString(a)
	c := simhash.Divide4(b)
	fmt.Println(b)
	fmt.Println(c)

	d := c[0] | c[1] | c[2] | c[3]
	e := simhash.HammingDistance(b, d)
	fmt.Println(e)
}
