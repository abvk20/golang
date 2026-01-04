package main

import "fmt"

func learnString() {
	fmt.Println("# STRINGS")
	fmt.Println("1. Interpreted Strings")
	fmt.Print("It will evaluate \n Then print\n")
	fmt.Println("2. Raw Strings")
	fmt.Println(`\nThis will not evaluate \n anything`)
	fmt.Println(`But it will print exactly what
is inside
these backticks.`)
	fmt.Println("\n---------Different Declarations in Go---------\n")

	var myString string = "Hello World\n"      // Declared and Initialized
	var myString2 = "Go will infer it\n"       // Inffered Declaration and initialization
	myString3 := "For Short hand := is used\n" // short declaration syntax
	// Go will declare a variable "myString3"
	fmt.Println(myString, myString2, myString3)
}
