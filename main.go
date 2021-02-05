package main

import (
	"fmt"
	"strconv"
)

func main() {
	a := "214748364799"
	b,_ := strconv.Atoi(a)
	fmt.Println(b)
}
