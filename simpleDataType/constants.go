package main

import (
	"fmt"
	"reflect"
)

func learnConstants() {
	// Implicitly Typed
	const a = 42
	fmt.Println("Constant is - a with value", a)
	fmt.Println("Constant is - a with type", reflect.TypeOf(a))
	fmt.Println(56 + a)
	fmt.Println(78.7 + a)

	// Explicitly typed Constant
	// Now it will only be used where there are string types
	const s string = "42"
	fmt.Println("Constant is - s with value", s)
	fmt.Println("Constant is - with type", reflect.TypeOf(s))
	fmt.Println("ehhlo" + s)

	// one const can assign to another
	const c = a

	// for multiple
	const (
		t   = s
		d   = c
		pie = 3.14
		y   = pie * d // This is calculated at COMPILE TIME
	)

}
