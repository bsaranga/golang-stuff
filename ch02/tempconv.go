// Package main is the L02 exercise: a temperature converter that exercises
// var / := / const, typed vs untyped constants, iota enumerations, and
// explicit numeric conversions.
//
// Reflection answers go here once you've finished (step 6):
//  1. <what was the compiler error when you mixed int32 y with an int literal,
//     and why do the untyped constants from step 2 not produce the same error?>
//  2. <what value does Kelvin end up with after the iota block, and why?>
package main

import "fmt"

// TODO (step 1): declare `type Scale int` and a const(...) block using iota
// to define Celsius, Fahrenheit, Kelvin as values of type Scale.

// TODO (step 2): declare AbsoluteZeroC, FreezingC, BoilingC as UNTYPED
// package-level constants. No type annotations.

// TODO (step 3): implement these three functions.
// func CtoF(c float64) float64 { ... }
// func CtoK(c float64) float64 { ... }
// func FtoC(f float64) float64 { ... }

// TODO (step 4): implement `describe`.
// func describe(s Scale, value float64) string {
//     // switch on s, return e.g. "100.00 °C"
// }

func main() {
	// TODO (step 5):
	//   - one `var x float64` with no initializer (print it to see the zero value).
	//   - one `:=` short declaration with an inferred type.
	//   - one `var y int32 = 42`; try `var z int = y + 1` first (it should fail
	//     to compile). Then fix it with an explicit conversion.
	//   - print describe(Fahrenheit, CtoF(BoilingC)) and describe(Kelvin, CtoK(AbsoluteZeroC)).

	fmt.Println("L02 — fill me in")
}
