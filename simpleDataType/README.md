# Simple Data Types 
These are data types that contain 1 value in it.
1. Strings
2. Numbers
3. Booleans
4. Errors

# Strings
There are 2 ways
- Interpreted String

"It uses normal quotes"
"It evaluates \n then outputs"

- Raw String

\`It uses backtick\`
\'It don't evaluate \n things \`

# Numbers
- int - integers
- uint - unsigned int
- float32
- float64
- complex64 - complex numbers
- complex128

# Booleans
- true
- false

# Error
It doesn't have exception mechanism like other languages.
In Go it is built in the language.
So in Go we simply return errors when things go wrong.
If value is `nil` it would mean no error.
---
type error interface {
    Error() string
}
---
The `Error` is a type
It is value that has a method   
So you can call a function on that Error
and
that function or method is the Error function
and
that's going to return the Error message

#Variables
Declared using syntax
`var` keyword `nameOfVar` `type - string, number, boolean, error, int, uint, float32, float64, complex64, complex128`

`var customName string`