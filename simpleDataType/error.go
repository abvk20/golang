package main

import "fmt"

func learnErrors() {
	fmt.Println(`
	It doesn't have exception mechanism like other languages.
In Go it is built in the language.
So in Go we simply return errors when things go wrong.
If value is nil it would mean no error.
---
type error interface {
    Error() string
}
---
The Error is a type
It is value that has a method   
So you can call a function on that Error
and
that function or method is the Error function
and
that's going to return the Error message`)
}
