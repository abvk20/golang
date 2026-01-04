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
}
