package main

import "fmt"

func main() {
	var username string = "Tushar"
	fmt.Println(username)
	fmt.Printf("The Variable type is: %T\n", username)

	var isLogin bool = true
	fmt.Println(isLogin)
	fmt.Printf("The Variable type is: %T\n", isLogin)

	var number int = 100
	fmt.Println(number)
	fmt.Printf("The Variable type is: %T\n", number)

	//alis method to delcare variable
	var name string
	fmt.Printf("The Variable type is: %T\n", name)

	//implict way to declare variable
	var number1 = 10
	fmt.Println(number1)

	//without var
	number3 := 123
	fmt.Println(number3)

	/*
      Note := is walrus operator and it is used only inside function not
	  outside the function and also not declear global...
	 */

}
