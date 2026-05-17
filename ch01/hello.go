// Package main is the entry point for L01.
//
// Starter shape: everything lives in one file. Your job (see README.md) is to
// extract the greeting into a sibling package `greet`, import it here, and
// have main do nothing but compose + print.
package main

import "fmt"

func main() {
	// TODO(L01): replace the inline greeting below with a call to
	//   greet.Greet("world")
	// from the package at ch01/greet. Import path:
	//   github.com/bsaranga/golang-stuff/ch01/greet
	msg := "hello, world"
	fmt.Println(msg)
}
