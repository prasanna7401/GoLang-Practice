package main

import (
	"fmt"
)

func main() {

	var m1 map[string]int
	// m1 := map[string]int // => ERROR
	fmt.Printf("%#v\n", m1)

	m2 := map[int]string{
		1: "One",
		2: "Two", // => Trailing comma is must for multi-line maps
	}

	m2[3] = "Three"
	fmt.Println(m2)

	delete(m2, 2)
	fmt.Println(m2)

	for k, v := range m2 {
		fmt.Printf("Key: %v, Value: %v\n", k, v)
	}
}
