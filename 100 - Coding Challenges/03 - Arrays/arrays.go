package main

import "fmt"

func main() {

	// BASICS

	var cities = [2]string{"New York", "Los Angeles"}
	fmt.Printf("%#v\n", cities) //=> [2]string{"New York", "Los Angeles"}

	grades := [3]float64{10.1, 12.3} // => [3]float64{10.1, 12.3, 0}
	fmt.Printf("%#v\n", grades)

	taskDone := [...]bool{}
	fmt.Printf("%#v\n", taskDone) // => [0]bool{}

	for _, city := range cities {
		fmt.Println(city)
	}

	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}

	nums := [...]int{30, -1, -6, 90, -6}
	positiveEvenCount := 0
	for _, num := range nums {
		if num > 0 && num%2 == 0 {
			positiveEvenCount++
		}
	}
	fmt.Println("Count of positive even numbers:", positiveEvenCount)
}
