package main

import (
	"fmt"
	"time"
)

func main() {

	// Using a for loop and an if statement print out all the numbers between 1 and 50 that divisible by 7.

	for i := 1; i <= 50; i++ {
		if i%7 == 0 {
			fmt.Println(i)
		}
	}

	// OR Change the code from the previous exercise and use the continue statement and use the break statement to print out only the first 3 numbers divisible by 7 between 1 and 50.
	count := 0
	for i := 1; i <= 50; i++ {
		if i%7 != 0 {
			continue
		}
		fmt.Println(i)
		count++
		if count == 3 { // stop after printing 3 numbers
			break
		}
	}

	// Using a for loop, an if statement and the logical and operator print out all the numbers between 1 and 500 that divisible both by 7 and 5.
	for i := 1; i <= 500; i++ {
		if i%7 == 0 && i%5 == 0 { // if i is divisible both by 7 and 5
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println("")

	currentYear := time.Now().Year()
	for birthYear := 2000; birthYear <= currentYear; birthYear++ {
		if birthYear%4 == 0 && (birthYear%100 != 0 || birthYear%400 == 0) {
			fmt.Printf("%d\n", birthYear)
		}
	}

	age := -9
	switch true {
	case age < 0 || age > 100:
		fmt.Println("Invalid age")
	case age < 18:
		fmt.Println("You are a minor")
	case age == 18:
		fmt.Println("Congrats! You are an adult now")
	default:
		fmt.Println("You are an adult")
	}
}
