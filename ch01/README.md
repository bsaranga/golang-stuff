# L01 — Packages, `main`, imports, `go run` vs `go build`

## Objective

Get comfortable with the smallest possible Go program shape, then prove you understand it by splitting a one-file hello into **two packages** in the same module.

You should finish this lesson able to answer, without looking it up:

1. Why must the entry-point package be named `main`, and why must it contain `func main()`?
2. What does `go run .` do that `go build` does not?
3. When you write `import "github.com/bsaranga/golang-stuff/ch01/greet"`, where does Go look for that code?
4. Why is the directory name *almost* the package name, but not strictly?
5. Why is `Greet` exported but `greet` is not?

## Key idioms / mental model

- **One directory = one package.** Every `.go` file in `ch01/greet/` must declare `package greet`. You cannot mix package names in a directory (test files aside).
- **`package main` is special.** It is the only package that produces an executable, and it must define `func main()`. Library packages have any other name.
- **Imports are by module path, not file path.** The string in `import "..."` is the module path from `go.mod` plus the subdirectory.
- **Exported = capitalized.** `Greet` is visible to importers; `greet` (lowercase) is package-private. There is no `public`/`private` keyword — capitalization *is* the access modifier.
- **`go run .`** compiles to a temp dir and runs. **`go build`** produces a binary in the current directory. **`go install`** puts it in `$GOBIN`. All three use the same compiler; the difference is where the output lands.
- **`go.mod` is the manifest.** `module <path>` declares the import prefix for this codebase. `go <version>` pins the language version.

## What's in this directory

- `hello.go` — a single-file program. Runs as-is. Your job is to refactor it.
- `greet/` — empty. You will create `greet/greet.go` here.

## Exercise

1. Run the starter as-is and confirm output:
   ```
   go run ./ch01
   ```
2. Create `ch01/greet/greet.go`. Declare `package greet`. Move the greeting logic into an **exported** function `Greet(name string) string` that returns (not prints) the greeting.
3. Edit `ch01/hello.go` to import `github.com/bsaranga/golang-stuff/ch01/greet` and call `greet.Greet("world")`, printing the result with `fmt.Println`.
4. Re-run: `go run ./ch01`. Output must be unchanged.
5. Build a binary: `go build -o /tmp/hello ./ch01 && /tmp/hello`. Confirm the same output.
6. **Stretch:** add an unexported helper `salutation()` inside `greet` that `Greet` calls. Try to call `greet.Salutation()` from `main` — observe and explain the compile error.

## Gotchas to watch for

- Forgetting `package greet` at the top of `greet/greet.go` — you'll get a confusing "expected 'package', found ..." error.
- Naming the function `greet` (lowercase) in the `greet` package — it compiles inside the package but `main` cannot see it. The error message ("cannot refer to unexported name") is the language teaching you about capitalization.
- Importing by directory path (`"./greet"`) — that's an old GOPATH-era pattern and will not work with modules. Always use the full module path.
- Two `.go` files in `ch01/greet/` with different `package` declarations — Go will refuse to compile the directory.

## Checkpoint

When you can run `go run ./ch01` from the repo root and the output is identical to the starter, but the greeting logic lives in `ch01/greet/greet.go` and is called via an exported function, you're done. Commit, then ping me for L02.
