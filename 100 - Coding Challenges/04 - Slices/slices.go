package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	num1 := []string{"A", "B", "C"}
	fmt.Println(num1)

	for _, v := range num1 {
		fmt.Println(v)
	}

	// Solve the Errors in the following code

	mySlice := []float64{1.2, 5.6}
	// mySlice[2] = 6 // => ERROR -> index out of range [2] with length 2

	a := float64(10) // ERROR -> cannot use a (type int) as type float64 in assignment
	mySlice[0] = a

	// mySlice[3] = 10.10 // => ERROR -> index out of range [2] with length 2

	mySlice = append(mySlice, a)
	fmt.Println(mySlice)

	// END

	// Append a slice to another
	nums := []float64{1.1, 2.2, 3.3}
	fmt.Println(nums)

	nums = append(nums, 10.1)
	nums = append(nums, 4.1, 5.5, 6.6)

	fmt.Println(nums)

	n := []float64{100.1, 200.2}

	nums = append(nums, n...)
	fmt.Println(nums)

	// END

	// Calculate the sum and product of command line arguments
	sum := 0
	product := 1
	args := os.Args[1:]

	if len(args) < 2 || len(args) > 10 {
		fmt.Println("Please enter between 2 and 10 numbers!")
	} else {
		for _, inputs := range args {
			numbers, err := strconv.Atoi(inputs) // Convert string user inputs to int
			if err == nil {
				sum += numbers
				product *= numbers
			}
		}
	}

	fmt.Printf("Sum: %d, Product: %d\n", sum, product) // => Sum: 10, Product: 30 (for inputs 2 3 5)

	// END

	// Calculate the sum of the middle elements
	someNumbers := []int{5, -1, 9, 10, 1100, 6, -1, 6}
	middleElements := someNumbers[2 : len(someNumbers)-2]
	addedValue := 0

	for _, value := range middleElements {
		addedValue += value
	}
	fmt.Printf("Sum of middle elements %v: %d\n", middleElements, addedValue) // => Sum of middle elements [9 10 1100 6]: 1125

	// END

	// Create a copy of the friends slice to prove they don't refer to same object
	friends := []string{"Marry", "John", "Paul", "Diana"}

	enemies := make([]string, len(friends))

	copy(enemies, friends)

	fmt.Printf("Friends: %v is stored in location %p\n", friends, &friends) // => Friends: [Marry John Paul Diana] is stored in location 0xc0000080c0
	fmt.Printf("Enemies: %v is stored in location %p\n", enemies, &enemies) // => Enemies: [Marry John Paul Diana] is stored in location 0xc0000080d8

	// END

	// newYears slice should contain first and last three years in list
	years := []int{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009, 2010}
	newYears := years[:3]
	newYears = append(newYears, years[len(years)-3:]...) // Appending the rest of the years

	fmt.Printf("%v\n", newYears) // => [2000 2001 2002 2008 2009 2010]
}
