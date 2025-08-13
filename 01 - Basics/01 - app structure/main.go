package main // default first clause in a Go file, indicating the package name

import "fmt" // import the fmt package for formatted I/O

const name = "Prasanna"

var age int = 25 // declare a variable with an initial value as integer

// main function is the entry point of the program
// does not take any arguments and does not return any value
// it is where the program starts executing
func main() {
	fmt.Println("Hello, World!")
	fmt.Printf("Your name is %s\n", name)
	fmt.Printf("Your age is %d\n", age)
}
