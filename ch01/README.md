# L01 — Packages, `main`, imports, `go run` vs `go build`

> **Retake iteration 3.** Prior attempts: `lesson_completion_1.md`, `lesson_completion_2.md`.
> Focus this pass: the **directory name ↔ `package` declaration ↔ import path ↔ caller-side identifier** chain. Specifically: aliases are *optional*, used only for caller-side disambiguation — not because dir and package name disagree. And the *why* behind capitalization-as-visibility (deeper than "neat").

## Objective

Same as before: get fluent with the smallest Go program shape and the package boundary. The objective hasn't moved — your understanding has, and this pass drills the one remaining soft spot from iterations 1–2.

You should finish this lesson able to answer, without looking it up:

1. Given a Go source file with `import "github.com/x/y/foo/bar"` at the top, which identifier do you write on the left-hand side of `.` at the call site — `foo`, `bar`, or something else? What determines it, and what *doesn't*?
2. When is an `import` alias (e.g. `b "github.com/x/y/foo/bar"`) actually *necessary* vs. merely stylistic? Give a concrete situation where the compiler will refuse to build without one.
3. Two `.go` files in the same directory must agree on their `package X` declaration. Why does the language enforce this — what would break in tooling, builds, or callers if Go allowed a directory to contain multiple packages?
4. Go uses capitalization to mark exported identifiers instead of a `private`/`public` keyword. What practical consequence does this have at the *call site* that a keyword-based system wouldn't give you? Answer in terms of what a human (or grep) reading code can do.
5. What's the difference between *the directory name*, *the package name declared inside the `.go` files*, and *the identifier you use in `import` paths*? Which two are coupled, and which one is free to differ?

## Key idioms / mental model

- **Four distinct things, often conflated:**
  1. **Directory name** on disk (e.g. `ch01/messages/`).
  2. **Package declaration** inside the `.go` files (`package greet`).
  3. **Import path** in callers (`"github.com/.../ch01/messages"`).
  4. **Caller-side identifier** before the dot (`greet.Something`).
  Mapping: (1) and (3) are coupled — the import path is the directory. (2) and (4) are coupled — the caller identifier is the declared package name. (1) and (2) are *independent*: convention says match them, but the compiler doesn't require it.
- **Aliases are for the caller, not for fixing mismatches.** `foo "github.com/.../bar"` says "in *this file*, call the bar package `foo`." You use it when two imports collide (e.g. two different packages both declared as `greet`), or when you want a shorter local name. You do **not** need an alias just because a directory and its package declaration disagree — the unaliased import already binds to the declared package name.
- **Capitalization as visibility, not just terseness.** A keyword-based system (`public func Foo`) puts the visibility info at the *declaration*. Go's capitalization rule puts it at *every reference*: `foo.Bar(x)` tells you at the call site, with no jump-to-definition, that `Bar` is exported. `grep -nE '[a-z]+\.[A-Z]'` finds every cross-package call in a tree.
- **One directory = one package.** Not a style rule — a *toolchain* invariant. Files in a directory are compiled together as a single linkage unit; mixed package declarations would make the unit ambiguous.
- **`go run` and `go build` both compile.** Neither catches errors the other misses.

## What's in this directory

- `hello.go` — `package main`. Imports nothing yet. Your job is to wire it up to packages you'll build.
- No subdirectories yet. You will create them.

## Exercise

You're building a tiny greetings program that prints two different greetings, sourced from two packages **whose declared names collide**. This will force you to confront the alias mechanism for its actual purpose.

1. Create `ch01/messages/messages.go`. **Constraint:** declare it as `package greet` (NOT `package messages`). Inside, export `Morning(name string) string` returning `"Good morning, <name>."`. Add an unexported helper that `Morning` calls.

2. Create `ch01/notes/notes.go`. **Constraint:** declare it *also* as `package greet`. Inside, export `Evening(name string) string` returning `"Good evening, <name>."`. Also use an unexported helper.

3. Edit `hello.go` to import both packages and print both greetings. You will hit a compile error if you try the naive import — diagnose it, then resolve it with an alias on **exactly one** of the imports (your choice which). Then call both exported functions and print the results on separate lines.

4. Once it builds: `go run ./ch01`. Output should be two lines, morning then evening.

5. **Reflection (write your answers as comments in `hello.go`):**
   - On line 1 of `hello.go`'s comment block: which directory does the *import path* for `messages.go` point to, and which identifier do you use at the call site? Are those the same word?
   - On line 2: if you removed the alias from your imports, exactly which compiler error would you get, and from which line?

6. **Stretch:** delete the unexported helper from one of the packages and try calling it as if it were exported from `hello.go`. Read the error. Then think: how would the same code look in a language with `public`/`private` keywords, and what does a reader lose?

## Gotchas to watch for

- "The directory is `messages/` but I declared `package greet` — do I need an alias to import it?" **No.** Try it without one first. The bare `import "github.com/.../ch01/messages"` already binds the local identifier to `greet`. An alias is only needed when something else — a *second* import — also wants that identifier.
- "I can put multiple packages in one directory if I declare them differently in each file." **No.** Go refuses to compile a directory with mixed `package X` declarations (test files in `package X_test` excepted). Step 1 and step 2 above use two *different* directories that happen to declare the same package name — that's legal and exactly the situation where aliases earn their keep.
- "Capitalization is just shorter than writing `public`." There's a real, observable difference at the call site. If you can't articulate it (Q4), you don't have the mental model yet — re-read the Key idioms bullet on this.
- `go vet ./...` will flag a name collision more clearly than `go build` sometimes. If the compiler error feels cryptic, run vet too.

## Checkpoint

You're done when:
- `go run ./ch01` prints both greetings.
- Exactly one of your imports uses an alias, and you can explain in one sentence why an alias was *necessary* there (not just convenient).
- You can name, without looking: which Go construct controls the import path, which controls the caller-side identifier, and which has neither role.

Then run `/go:complete 01` again.
