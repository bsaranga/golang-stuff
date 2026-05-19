# Lesson 01 — Completion (Retake)

**Date:** 2026-05-18
**Final grade:** B (17/20)
**Iteration:** 2 (prior: `lesson_completion_1.md`)

## Q1. Trace what Go does between reading `import "github.com/bsaranga/golang-stuff/ch01/farewell"` and opening files on disk — which file does it consult first, what does it strip, and how does it find the files?
**Answer:** First the compiler will look for the go.mod file which contains the module declaration, this is the prefix, and the prefix is stripped from the import string and looks from the root to find the files, and compiles every go file to a package
**Rating:** 4/4 — Full trace: `go.mod` → module prefix → strip → relative-to-module-root directory → all `.go` files in it form the package. Direct improvement from iteration 1's 2/4.
**Correct take:** `go.mod`'s `module` line gives the prefix. The toolchain strips that prefix from the import path; the remainder is a directory relative to the module root. Every `.go` file in that directory is compiled as one package.

## Q2. If `greet/greet.go` had stayed as `package greetings`, what would `hello.go` need to look like to call `Greet`? Be specific about the import line and the call site.
**Answer:** If the package name doesn't match the directory name here's how I'd import it, `greetings "github.com/bsaranga/golang-stuff/ch01/greet"` where greetings is the package name in the greet directory. This way I can have multiple packages in one directory too. :)
**Rating:** 2/4 — Reaches for aliasing, but (a) the alias `greetings` is redundant — an unaliased `import "...ch01/greet"` already binds the identifier to `greetings`, the package's declared name; (b) "multiple packages in one directory" is incorrect — one directory = one package (test files aside); (c) the call site (`greetings.Greet(...)`) wasn't addressed.
**Correct take:** Just `import "github.com/bsaranga/golang-stuff/ch01/greet"` — no alias needed — and call `greetings.Greet("world")`. The import path locates the directory; the package's own `package X` declaration determines the caller-side identifier. Aliases are only for renaming when *the caller* wants something different (e.g. to disambiguate two imports with the same package name).

## Q3. Name one situation where only `go build` will do, and one where `go run` is clearly the better fit. (No "build catches errors.")
**Answer:** I'd use `go build` when I need an artifact that I want to share with someone or if I want to observe the artifact or benchmark it against something. However during normal development I'd use `go run` as it compiles to a temp directory and runs immediately and discards the temp artifact without me having to manage it, this is a much faster iteration during development when I just want to check the program result.
**Rating:** 4/4 — Clean artifact-vs-ephemeral framing, with concrete scenarios on both sides. Direct improvement from iteration 1's 2/4.
**Correct take:** `go build` when you need a reusable artifact — shipping, deploying, benchmarking, repeated invocation without recompiling. `go run` for the tight edit-run loop where the artifact is noise.

## Q4. What happens if `hello.go` calls `farewell.sayBye("world")` directly? What's the error, and why is the rule tied to capitalization rather than a `private` keyword?
**Answer:** I'd get the error, `name sayBye not exported from the package farewell`. Capitalization is a simpler way of denoting a method as public/private without an additional keyword. I guess the language designers thought its neat.
**Rating:** 3/4 — Error and rule are right; the *why* is hand-wavy ("they thought it's neat").
**Correct take:** Compile error: `cannot refer to unexported name farewell.sayBye`. The deeper why: visibility is readable at the call site itself, without needing to find the declaration. `farewell.sayBye` *looks* internal from where you read it — no jumping to definitions, no IDE help required. It's part of Go's broader bias toward source-level transparency.

## Q5. Someone adds `ch01/greet/extra.go` declaring `package extras`. What happens, at what stage, and does it matter if `extra.go` is actually used?
**Answer:** I would get a build error saying that there are two packages in one directory. Even if anything in extra.go is actually used in my program, the error would still happen. I think one package = one directory.
**Rating:** 4/4 — Correct at every level: compile-time failure, package-discovery stage, irrelevant whether anything in `extra.go` is consumed.
**Correct take:** Compile error at the *package-loading* stage — `found packages greet (greet.go) and extras (extra.go) in ...ch01/greet`. The toolchain refuses to compile a directory with mixed package declarations regardless of what's used; one directory = one package (test files in `_test.go` with `package X_test` excepted).

## Overall feedback
Big jump from iteration 1 on the two target weak areas: the `go.mod` → on-disk resolution trace (Q1) and the `go run` vs `go build` distinction (Q3) are now clean. The remaining soft spot is the *directory ↔ package declaration ↔ import identifier ↔ caller identifier* chain (Q2) — the misconception that aliasing is *required* when names diverge, plus the stray belief that multiple packages can coexist in one directory (which Q5 then correctly contradicts — worth reconciling those two intuitions). For the "why" behind capitalization (Q4), reach for *source-level transparency* rather than "neat." Revisit the **Key idioms** bullet on "Directory name vs package name are independent" in `ch01/README.md` once more.
