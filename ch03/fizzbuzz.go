// Package main is the L03 exercise: FizzBuzz (twice) plus the Sieve of
// Eratosthenes, to drill if / for / switch and the labeled-break / init-clause
// idioms.
//
// Reflection answers go here once you've finished (step 5):
//  1. <why does case ordering matter in fizzbuzzSwitch?>
//  2. <why does the sieve's inner loop start at i*i instead of 2*i?>
//  3. <could the outer sieve loop be a `for i := range ...`? why or why not?>
package main

import "fmt"

// fizzbuzzIf prints FizzBuzz for 1..n using if / else if / else.
//
// TODO (step 1): implement using a classic three-part for loop and
// if / else if / else. Order the conditions so the % 15 case wins.
func fizzbuzzIf(n int) {
	// TODO
}

// fizzbuzzSwitch prints FizzBuzz for 1..n using a tagless switch.
//
// TODO (step 2): implement using `switch { case ...: ... }` with no tag.
// The case order matters — think about why before you write it.
func fizzbuzzSwitch(n int) {
	// TODO
}

// primesUpTo prints every prime p with 2 <= p <= n, using the Sieve of
// Eratosthenes.
//
// TODO (step 3): implement. Use a []bool of length n+1, where index i means
// "i has been marked composite". For each i from 2 while i*i <= n, if i is
// not composite, mark every multiple of i from i*i upward. Then walk the
// slice and print every index that is still unmarked. Use the condition-only
// form of `for` (`for cond { ... }`) somewhere it reads naturally.
func primesUpTo(n int) {
	// TODO
}

// findFirstFactor returns the smallest prime factor of n. Stretch goal — see
// step 6 in the README. Use an `if` with an init clause inside the loop body.
func findFirstFactor(n int) int {
	// TODO (stretch)
	return 0
}

func main() {
	// TODO (step 4):
	//   fizzbuzzIf(15)
	//   fmt.Println("---")
	//   fizzbuzzSwitch(15)
	//   fmt.Println("---")
	//   primesUpTo(50)
	_ = fmt.Println // remove this once you've written real Println calls
}
