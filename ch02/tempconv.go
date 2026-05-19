// Package main is the L02 exercise: a temperature converter that exercises
// var / := / const, typed vs untyped constants, iota enumerations, and
// explicit numeric conversions.
//
// Reflection answers go here once you've finished (step 6):
//  1. <what was the compiler error when you mixed int32 y with an int literal,
//     and why do the untyped constants from step 2 not produce the same error?>
//     Answer 1: the compiler error is "cannot use yInt (type int32) as type int in addition", the same arithmetic works with untyped constants because they can be implicitly converted to the appropriate type.
//  2. <what value does Kelvin end up with after the iota block, and why?>
//     Answer 2: Kelvin ends up with the value 2 since iota starts from 0 and increments by 1 on each line within the const block.
package main

import "fmt"

// TODO (step 1): declare `type Scale int` and a const(...) block using iota
// to define Celsius, Fahrenheit, Kelvin as values of type Scale.
type Scale int

const (
	Celsius Scale = iota
	Fahrenheit
	Kelvin
	Rankine
)

// TODO (step 2): declare AbsoluteZeroC, FreezingC, BoilingC as UNTYPED
// package-level constants. No type annotations.
const AbsoluteZeroC = -273.15
const FreezingC = 0.0
const BoilingC = 100.0 // reflection: upon typing this to float32, the compiler gave an error saying that the function expects a float64 argument, since Go does not implicitly convert between different numeric types, and the untyped constant can be used as a float64 without issue.

// TODO (step 3): implement these three functions.
// func CtoF(c float64) float64 { ... }
// func CtoK(c float64) float64 { ... }
// func FtoC(f float64) float64 { ... }
func CtoF(c float64) float64 {
	return c*9.0/5.0 + 32.0
}

func CtoK(c float64) float64 {
	return c + (-AbsoluteZeroC)
}

func FtoC(f float64) float64 {
	return (f - 32.0) * 5.0 / 9.0
}

// TODO (step 4): implement `describe`.
//
//	func describe(s Scale, value float64) string {
//	    // switch on s, return e.g. "100.00 °C"
//	}
func describe(s Scale, value float64) string {
	switch s {
	case Celsius:
		return fmt.Sprintf("%.2f °C", value)
	case Fahrenheit:
		return fmt.Sprintf("%.2f °F", value)
	case Kelvin:
		return fmt.Sprintf("%.2f K", value)
	default:
		return fmt.Sprintf("%.2f (unknown scale)", value)
	}
}

func main() {
	// TODO (step 5):
	//   - one `var x float64` with no initializer (print it to see the zero value).
	//   - one `:=` short declaration with an inferred type.
	//   - one `var y int32 = 42`; try `var z int = y + 1` first (it should fail
	//     to compile). Then fix it with an explicit conversion.
	//   - print describe(Fahrenheit, CtoF(BoilingC)) and describe(Kelvin, CtoK(AbsoluteZeroC)).
	var x float64
	fmt.Println(x)
	y := 3.14
	fmt.Println(y)
	var yInt int32 = 42
	var z int = int(yInt) + 1
	fmt.Println(z)
	fmt.Println(describe(Fahrenheit, CtoF(BoilingC)))
	fmt.Println(describe(Kelvin, CtoK(AbsoluteZeroC)))
}
