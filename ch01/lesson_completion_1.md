# Lesson 01 — Completion

**Date:** 2026-05-18
**Final grade:** B (15/20)

## Q1. When you write `import "github.com/bsaranga/golang-stuff/ch01/greet"` in `hello.go`, how does Go actually resolve that string to the files on disk? Walk me through what it uses to find them.
**Answer:** It uses the go.mod or the module file to find them
**Rating:** 2/4 — Names the right file but skips the actual resolution mechanism.
**Correct take:** Go reads the `module` line in `go.mod` (`github.com/bsaranga/golang-stuff`). It strips that prefix from the import path, treats the remainder (`ch01/greet`) as a directory relative to the module root, and compiles every `.go` file in that directory as one package.

## Q2. Why must the entry-point package specifically be named `main`, and why must it contain `func main()`? What would happen if you named the package `app` instead but still defined `func main()` inside?
**Answer:** Only the main package can be compiled to a binary, and the main() function is what executes when the binary is run
**Rating:** 3/4 — Core is right; didn't address the `app` variant explicitly.
**Correct take:** `package main` is the linker's signal to produce an executable; any other package name is treated as a library and won't link an entry point even if a `func main()` exists inside. The toolchain ties "executable" to the package name, not the function name.

## Q3. Suppose you renamed `salutation` to `Salutation` (capitalized) but kept it only used internally by `Greet`. Does anything functionally change? Is there a reason you'd prefer to keep it lowercase?
**Answer:** Nothing changes functionally but now `Salutation` is an exported function and it can be used in another go package, since this method is meant to be internal its best to keep it lower case which would be unexported.
**Rating:** 4/4 — Correct and shows the encapsulation instinct.
**Correct take:** Same runtime behavior, but exporting it widens the package's API surface and ties you to backward compatibility. Keep internal helpers lowercase so the package boundary stays honest.

## Q4. Beyond "where the binary lands," is there any practical reason you'd reach for `go build` over `go run` during normal development? When does each one fit?
**Answer:** I'd run go build to see if the binary gets produced and I do not run into any compilation errors, but go run is preferred since it shows program output immediately
**Rating:** 2/4 — Mixes up the compile-check angle; `go run` also compiles.
**Correct take:** Both compile, so neither is "the" compile check (`go build` without `-o` and `go vet`/`go test` are the usual checks). The real split: `go run` is for tight edit-run loops where you don't want an artifact; `go build` is when you need a reusable binary — to ship, to benchmark, to run repeatedly without recompiling, or to inspect with tools.

## Q5. Imagine you delete `func main()` from `hello.go` but leave `package main` at the top. What does `go run ./ch01` do? What about `go build ./ch01`?
**Answer:** I'd get a compilation error saying that main() is undeclared in the main package, same with go build.
**Rating:** 4/4 — Correct; both fail at link time with the missing-`main` error.
**Correct take:** Both fail identically — the linker requires `func main()` in `package main`. The error is `function main is undeclared in the main package`.

## Overall feedback
Mental model for exports, package boundaries, and the `main`/`func main()` contract is solid — Q3 and Q5 show you've internalized the lesson's core. Two areas to firm up: the precise mechanics of import resolution via `go.mod` (Q1), and the real distinction between `go run` and `go build` — both compile, so the difference is artifact vs. ephemeral run, not "build catches errors." Re-skim the "Key idioms" section of `ch01/README.md` on imports and the three build commands before L02.
