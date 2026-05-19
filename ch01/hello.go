// Package main is the entry point for L01 (retake iteration 3).
//
// The exercise (see README.md) is to build two sibling packages — one under
// ch01/messages/, one under ch01/notes/ — that BOTH declare `package greet`,
// and import them both into this file. That collision is the whole point:
// it forces you to use an import alias for its actual purpose (caller-side
// disambiguation), not as a workaround for a directory/package-name mismatch.
//
// Reflection answers go in this comment block once you're done (step 5):
//  1. <which dir does the messages.go import path point to, and which
//     identifier do you use at the call site? are those the same word?>
//  2. <which compiler error would you get if you removed the alias, and
//     from which line?>
package main

import (
	"fmt"

	greet "github.com/bsaranga/golang-stuff/ch01/messages" // reflection: import path points to "messages" dir but we use the greet identifier (package name) at the call site.
	"github.com/bsaranga/golang-stuff/ch01/notes"          // reflection: i don't get any compiler errors unfortunately, my editor forces aliases on both imports
)

func main() {
	fmt.Println(greet.Morning("Saranga"))
	fmt.Println(notes.Evening("Saranga"))
}
