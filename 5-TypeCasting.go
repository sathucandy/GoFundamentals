package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 42
	fmt.Printf("%v, %T\n", i, i)

	var j float32
	j = float32(i) ////////HERE WE DO THE TPECASTING
	fmt.Printf("%v, %T\n", j, j)

	///CONVERTING INTEGER TO STRING
	var k string
	//	k = string(i) ////////THIS WONT WORK AS WE WILL GET UNICODE CHARACTER OF 42 WHICH IS ('*')
	k = strconv.Itoa(i) ///THIS WILL WORK
	fmt.Printf("%v, %T\n", k, k)
}
