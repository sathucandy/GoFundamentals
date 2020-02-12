package main

import "fmt"

// func add(x int, y int) int {
// 	out := x + y
// 	return out
// }

/////////////OR///////////
// func add(x, y int) int {
// 	out := x + y
// 	return out
// }

////////////OUTPUTTING MORE THAN ONE VALUES////////

// func calc(x int, y int) (int, int) { //////////////HERE WE ARE USING INT 2 TIMES AS ITS THE RETURN TYPE OF TWO OUTPUTS
// 	out1 := x + y
// 	out2 := x - y
// 	return out1, out2
// }

/////////////////////WE CAN ALSO WRITE IN TIS WAY TO RETURN THE VALUES

func calc(x int, y int) (out1, out2 int) { //////////////HERE WE ARE USING OUT1, OUT2 AS ITS THE RETURN TYPE OF TWO OUTPUTS
	out1 = x + y
	out2 = x - y
	return
}

func main() {
	num1 := 5
	num2 := 4
	result1, result2 := calc(num1, num2)
	fmt.Print(result1, result2)
}
