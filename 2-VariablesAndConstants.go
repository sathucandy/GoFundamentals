package main

import "fmt"

var I int = 42 ///DECLARING ANY VARIABLE IN CAPITAL WILL MAKE IT SCOPE PUBLIC

func main() {

	//DECLARING VARIABLES
	var num = 5
	var num2 int = 8
	var restult = num + num2

	//DECLARING VARIABLES IN ONE LINE
	var num5, num6 int
	num5, num6 = 2, 3
	fmt.Println(num5)
	fmt.Println(num6)

	var num3 int /// BY DEFAULT THE VALUE IS 0
	num3 = 5
	fmt.Println(num3)

	//fmt.Print(2 + 3)
	// fmt.Print(num)
	fmt.Println(restult)

	///////////IN THE ABOVE EXAMPLES WE ARE CREATING AND THEN ASSIGNING THEM VARIABLES
	///////////SO HERE IS ANOTHER WAY

	result2 := 10 ////HERE WE USE ':='     IT IS SAME AS VAR RESULT INT = 9;
	fmt.Println(result2)

	////CREATING A CONSTANT
	const abc = 9
	fmt.Println(num)
}
