package main
import("fmt")

func main(){
	var greeting string = "Hello, Go. You're welcome"

	fmt.PrintIn(greeting)
}


package main

import ("fmt")

func main()  {
	var (
	name string = "Benson"
	age int = 100
	money float32 = 2.90
	)

	fmt.Println("My name is " + name + ". I am",age, "years old. This is a", money,"float.")
}

// I want to comment

/* I want a multi-line
comment to be here
anyways
*/