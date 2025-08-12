package main

import (
	"fmt"
	"strconv"
	// "math" --> UNUSED
)

// ERROR -> := is not allowed in package scope (only local scope)
// version := "3.1"

var version = "3.1"

func main() {

	// =========================
	// VARIABLES
	// =========================

	var a float64 = 7.1

	x, y := true, 3.7

	// ERROR -> no new variables on left side of :=
	// a, x := 5.5, false

	a, x = 5.5, false

	_, _, _ = a, x, y

	// ERROR -> A string is initialized only using double quotes ("")
	// name = 'Golang'

	name := "Golang"
	fmt.Println(name)

	_ = version

	// =========================
	// CONSTANTS
	// =========================

	const (
		daysWeek   = 7
		lightSpeed = 299792458
		pi         = 3.14159
		secPerDay  = 24 * 60 * 60
		daysYear   = 365
	)

	fmt.Printf("%d\n", secPerDay*daysYear) // => 31536000

	//** There are ONLY boolean constants, rune constants, integer constants, **//
	//** floating-point constants, complex constants, and string constants. **//

	// declaring a constant of type slice int ([]int)
	// ERROR -> const initializer []int literal is not a constant
	// const m = []int{1: 3, 4: 5, 6: 8}	-> You cannot declare a slice constant.
	// _ = m

	// ERROR -> invalid operation: a * b (mismatched types int and float64)
	// const a int = 7
	const i = 7 // make `a` untyped constant
	const j float64 = 5.6
	const k = i * j

	s := 8
	_ = s
	// ERRROR ->  You cannot initiate a constant at runtime (constants belong to compile-time)
	// const xc int = x  // variables belong to runtime

	// ERROR ->  You cannot initiate a constant at runtime (constants belong to compile-time)
	// const xc int = x  // variables belong to runtime

	// ERROR ->  You cannot initiate a constant at runtime (constants belong to compile-time)
	// const noIPv6 = math.Pow(2, 128) // functions calls belong to runtime

	const (
		Jun = (iota + 6)
		Jul
		Aug
	)

	fmt.Println(Jun, Jul, Aug) // => 6 7 8

	// =========================
	// FMT
	// =========================

	m, n, o := 10, 15.5, "Gophers"
	score := []int{10, 20, 30}

	fmt.Printf("%d %.1f %s\n", m, n, o)         // => 10 15.5 Gophers
	fmt.Printf("%q\n", o)                       // => "Gophers"
	fmt.Printf("%v %v %v %v\n", m, n, o, score) // => 10 15.5 Gophers [10 20 30]
	fmt.Printf("%T %T\n", y, score)             // => float64 []int

	const g float64 = 1.422349587101
	fmt.Printf("%.4f\n", g) // => 1.4223

	shape := "circle"
	radius := 3.2
	circumference := 2 * pi * radius

	fmt.Printf("Shape: %q\n", shape) // => Shape: "circle"
	fmt.Printf("Circle's circumference with radius %f is %f\n", radius, circumference)
	// _ = shape

	// =========================
	// TYPE CONVERSION
	// =========================

	var e int = 3
	var f float64 = 3.2

	// e = float64(e) => Error because it is a typed variable
	// f = int(f) => Error

	e = int(f)     // f (float64) to e (int)
	f = float64(e) // e (int) to f (float64)

	// make e & f as untyped variables

	fmt.Printf("%T %T\n", e, f)

	var p = 3
	var q = 3.2
	var s1, s2 = "3.14", "5"

	// p1 := string(p) // int to string - WRONG => p1:: string
	// OR
	p1 := strconv.Itoa(p) // int to string - BEST => p1:3: string

	// s2_new, err := strconv.ParseInt(s2, 10, 64) // string to int64, base 10 => 5: int64
	s2_new, err := strconv.Atoi(s2) // => 5: int
	_ = err
	_ = s1

	q1 := fmt.Sprintf("%f", q)            // float64 to string => 3.200000: string
	q2, err := strconv.ParseFloat(s1, 64) // string float64 to float64 => 3.14: float64
	_ = err

	fmt.Printf("p1:%v: %T,\ns2:%v: %T,\nq1:%v: %T,\nq2:%v: %T\n", p1, p1, s2_new, s2_new, q1, q1, q2, q2) // => p1: string, s2: int64

	const distanceFromSun = 149597870700

	timeForSunlightToEarth := distanceFromSun / lightSpeed
	fmt.Printf("%v seconds\n", timeForSunlightToEarth)

	// =========================
	// NAMED TYPES & ALIASES
	// =========================

	type duration int
	var hour duration

	fmt.Printf("hour value: %v, type: %T\n", hour, hour) // => hour value: 0, type: main.duration
	hour = 3600
	fmt.Printf("hour value: %v, type: %T\n", hour, hour) // => hour value: 3600, type: main.duration

	type mile float64
	type kilometer float64

	const m2km = 1.60934

	var mileBerlinToParis mile = 655.3
	var kmBerlinToParis kilometer = kilometer(mileBerlinToParis * m2km)

	fmt.Printf("Distance from Berlin to Paris: %v miles = %v kilometers\n", mileBerlinToParis, kmBerlinToParis)
}
