# L03 — Control flow

## Objective

Get fluent with Go's three control-flow primitives — `if`, `for`, `switch` — and the deliberate choices Go made when paring them down. Go has one loop keyword, no `while`, no `do…while`, no implicit `switch` fallthrough, and no parentheses around conditions. By the end of this lesson you should be able to *defend* those choices, not just live with them.

You should finish this lesson able to answer, without looking it up:

1. `if` supports an optional **init clause** (`if x := f(); x > 0 { ... }`). What pattern does this enable that a plain `if cond { ... }` can't express as cleanly? Where, exactly, is `x` in scope — including across `else if` and `else`?
2. Go has no `while`, no `do…while`, no `loop`. The single `for` keyword takes four shapes. Name them, and explain why collapsing all looping into one keyword is a *design choice*, not a limitation.
3. In C, `switch` cases fall through by default — you write `break` to stop. Go reverses this: each case is implicitly terminated. Why was the default flipped, and what does writing `fallthrough` cost a reader when you genuinely need it?
4. Inside a `for` loop with a `switch` nested inside, what does a bare `break` do? How do you break out of the surrounding `for` from inside the switch? Be precise.
5. A `switch` with **no tag** (e.g. `switch { case x < 0: ...; case x == 0: ...; default: ... }`) is idiomatic Go where C or Java code would use `if / else if / else`. What does the switch form communicate that the if-chain doesn't?
6. `for i, v := range coll` evaluates the range expression **once** before the loop starts. Why does that matter, and what bug does it prevent that a naive desugaring would introduce?

## Key idioms / mental model

- **`if` takes no parentheses around the condition.** The braces are mandatory; the parens are forbidden. This isn't cosmetic — it's part of why Go's parser doesn't need a semicolon-insertion hack the way C-family languages do. The init clause is what makes `if` actually pull its weight in Go: it lets you confine a short-lived variable (often an error) to exactly the scope where it's checked.

  ```go
  if n, err := strconv.Atoi(s); err != nil {
      return err
  } else {
      use(n)   // n is still in scope here
  }
  // n and err are gone here — they belonged to the if
  ```

- **`for` is Go's only loop.** Four shapes, in order from most explicit to least:

  ```go
  for i := 0; i < n; i++ { ... }   // classic three-part
  for cond { ... }                  // while-equivalent
  for { ... }                       // infinite — loop control via break/return
  for i, v := range coll { ... }    // range over array/slice/map/string/chan
  ```

  One keyword. Fewer concepts to teach. When you read Go code, you never have to ask "which loop construct did they pick?" — there's only one.

- **`switch` does not fall through.** Each `case` is implicitly its own block; control exits the switch when the case ends. If you actually want fallthrough — rare — you write the `fallthrough` keyword as the last statement of the case. The default is the safe behavior; the dangerous one is opt-in.

  ```go
  switch http.StatusOK {
  case 200, 204:    // multiple values, no fallthrough needed
      return "ok"
  case 301, 302:
      return "redirect"
  default:
      return "other"
  }
  ```

- **Tagless `switch` replaces `if / else if / else` chains.** When the cases each have their own boolean expression, `switch { case a < 0: ...; case a == 0: ... }` reads better than the equivalent if-chain: it signals "these are sibling branches of one decision," and the cases line up visually.

  ```go
  switch {
  case score >= 90:
      grade = "A"
  case score >= 80:
      grade = "B"
  default:
      grade = "F"
  }
  ```

- **`break` and `continue` apply to the innermost enclosing loop or switch.** Inside a `for` containing a `switch`, a bare `break` only exits the switch. To break the *loop* from inside the switch, label the loop:

  ```go
  outer:
  for _, row := range rows {
      switch row.kind {
      case "stop":
          break outer       // exits the for, not just the switch
      case "skip":
          continue outer    // skips to the next iteration of the for
      }
  }
  ```

- **`range` evaluates its expression once.** Before the loop body runs, the slice (or map, etc.) is captured. Mutating the underlying collection inside the loop does *not* change what `range` iterates over. This prevents a class of bug familiar from other languages where appending during iteration goes infinite or skips elements.

- **No ternary.** `x := cond ? a : b` doesn't exist in Go. Use a 4-line `if` instead. The language considers the visual cost worth paying for the simpler grammar.

## What's in this directory

- `fizzbuzz.go` — `package main`. Skeleton for the classic FizzBuzz over 1..100 plus a second function that prints primes up to N using the Sieve of Eratosthenes. Most logic is `// TODO`.

## Exercise

Two small programs in one file. You'll write FizzBuzz twice (once with `if/else if/else`, once with a tagless `switch`) and the prime sieve once — to feel where each control-flow shape earns its keep.

1. Implement `fizzbuzzIf(n int)` that prints, for each `i` from 1 to `n`:
   - `"FizzBuzz"` if `i` is divisible by 15,
   - `"Fizz"` if divisible by 3,
   - `"Buzz"` if divisible by 5,
   - otherwise the number itself.

   Use `if / else if / else` and `fmt.Println`. Classic three-part `for` loop.

2. Implement `fizzbuzzSwitch(n int)` with the same behavior, but use a **tagless `switch`** (no expression after `switch`). The case ordering matters — think about why.

3. Implement `primesUpTo(n int)` using the Sieve of Eratosthenes:
   - Create a `[]bool` of length `n+1`. Treat index `i` as "is `i` composite?" (zero value `false` = "not yet marked composite").
   - For each `i` from 2 to `√n`, if `i` is not composite, mark every multiple of `i` from `i*i` upward as composite.
   - After the sieve, walk the slice from 2 to `n` and print every `i` that is *not* marked composite.
   - Use the *condition-only* form of `for` (`for cond { ... }`) for at least one of the loops — pick where it reads most naturally.

4. In `main`, call `fizzbuzzIf(15)`, `fizzbuzzSwitch(15)`, and `primesUpTo(50)`. Print a separator line between each.

5. **Reflection (write your answers as comments at the top of `fizzbuzz.go`):**
   - Line 1: in `fizzbuzzSwitch`, why does the order of cases matter? What goes wrong if you put the `% 3` case before the `% 15` case?
   - Line 2: in `primesUpTo`, the inner marking loop starts at `i*i`, not at `2*i`. Why is `i*i` correct, and what does starting at `2*i` cost you?
   - Line 3: could you write the sieve's outer loop as `for i := range someSlice`? Why or why not?

6. **Stretch:** Add a `findFirstFactor(n int) int` that returns the smallest prime factor of `n`, using a `for` with an `if` containing an init clause to compute the factor and check it in one expression. Call it from `main` on a few numbers. The point is to force the init-clause shape — don't fall back to a precomputed variable.

## Gotchas to watch for

- **"Why won't `if (x > 0)` compile?"** Go forbids parens around the condition. The compiler error mentions a syntax issue; the fix is to drop the parens. Braces are required even for one-statement bodies — no `if x > 0 doThing()`.
- **The init-clause variable is gone after the `if`.** This is a feature, not a footgun, but worth internalizing: `if n, err := f(); err != nil { ... } else { use(n) }` — `n` is in scope in both branches, but *not* on the line after the closing brace.

  ```go
  if v := compute(); v > 0 {
      use(v)
  }
  // v is out of scope here — compile error if you try to use it
  ```

- **Bare `break` inside a `for` that contains a `switch` only exits the switch.** Easy to assume otherwise if you're coming from a language where `switch` is rarer. Labels exist for exactly this case.
- **`fallthrough` is rare and loud.** If you reach for it, ask whether listing multiple values on one case (`case 1, 2, 3:`) gets you what you want instead. Real fallthrough — where the *body* of case A should run, then continue into the body of case B — is the rare case the keyword was reserved for.
- **`for i := 0; i < len(s); i++` over a string gives you bytes, not characters.** `for i, r := range s` gives runes (with `i` as byte offset, not rune index). L02's aside `aside-bytes-and-unicode.md` covers the underlying reason.
- **Off-by-one in the sieve.** `[]bool` of length `n+1` so index `n` is valid. The outer loop runs while `i*i <= n` (not `i <= n`); the inner marks multiples *of* `i` starting at `i*i`.
- **`break` in a `for-range` doesn't unfreeze the range expression.** Whatever was captured at loop entry stays captured. There's no way to "rebind" the range mid-loop — you have to break out and start a new loop.

## Checkpoint

You're done when:
- `go run ./ch03` prints FizzBuzz twice (identical output from the two implementations) followed by all primes ≤ 50.
- You can articulate, without looking, why Go made `switch` non-falling-through by default and what scope rule applies to an `if`'s init variable across `else if` / `else`.
- You can name the four shapes of `for` and pick which one fits a given problem without thinking about it.
- The three reflection comments at the top of `fizzbuzz.go` are written and accurate.

Then run `/go:complete 03`.
