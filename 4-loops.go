package main

import "fmt"

//////////////IN GO LANGUAGE WE ONLY HAVE ONE LOOP WHICH IS FOR LOOP/////////////////////

func main() {
	// for {
	// 	//	fmt.Print("Hello World") /////////////////////SYNTAX1////////
	// }

	i := 1
	for i <= 5 {
		fmt.Println("Hello"+"World", i)
		i++
	}

	for i := 1; i <= 5; i++ {
		fmt.Println("Hello"+"World", i)
	}
}
