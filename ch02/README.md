# L02 — Variables, constants, basic types

## Objective

Get fluent with Go's declaration mechanics (`var`, `:=`, `const`, `iota`) and the *strictness* the language demands around types. By the end you should know which declaration form belongs in which context, why Go refuses implicit numeric conversions, and what "untyped constant" actually means at the call site.

You should finish this lesson able to answer, without looking it up:

1. Go has three declaration forms — `var x T = v`, `var x = v`, and `x := v` — plus `var x T` (no initializer). What does each say that the others don't? Which one is illegal at the package level, and why is that restriction useful rather than annoying?
2. Why does Go refuse `var x int = 1; var y float64 = x + 1.5`, but happily accept `const c = 1; var y float64 = c + 1.5`? Answer in terms of *when* the type of each value is fixed.
3. `byte` is an alias for `uint8`; `rune` is an alias for `int32`. The compiler treats `byte` and `uint8` as the same type — there's no conversion needed between them. Given that, why does the language bother giving them separate names? What does the choice between `byte` and `uint8` (or `rune` and `int32`) communicate to a reader?
4. Constants in Go can only hold values the compiler can evaluate at compile time — not arbitrary runtime expressions. List what *is* allowed, and explain the consequence: there's no way in Go to declare a runtime-computed *immutable* value. What's the idiomatic workaround?
5. `iota` resets to 0 at the start of each `const ( ... )` block and increments per line. Given `const ( A = iota; B; C )`, what are A, B, C? Now: how would you use `iota` to define bit-flag constants like `1, 2, 4, 8, ...`? Why does that pattern work?
6. What does Go do if you declare a local variable and never read it? What about an unread `const`? Why does the language treat these two cases differently?

## Key idioms / mental model

- **Three reasons to pick `var` over `:=`** inside a function: (a) you want the zero value (`var x int` — no `:=` form expresses this cleanly); (b) you're declaring at package level (`:=` is a function-only operator); (c) the literal's default type isn't the type you want (`var x byte = 20` reads better than `x := byte(20)`). Otherwise, `:=` is idiomatic inside functions.

  ```go
  var buf bytes.Buffer        // zero value — no := form expresses this
  var port uint16 = 8080      // overriding the default int type of 8080
  name := "claude"            // idiomatic default inside functions
  ```

- **Untyped constants are the escape valve.** Go is famously strict about numeric types — `int` and `int64` won't mix without an explicit conversion. *Constants* sidestep this because an untyped constant has no type until used in a context that demands one. That's why `const Pi = 3.14159` can be assigned to a `float32`, a `float64`, or used in an `int` expression (if it fits). The moment you write `const Pi float64 = 3.14159` you've thrown that flexibility away — use untyped constants by default; type them only when you specifically need the constraint.

  ```go
  const Tau = 6.28318          // untyped — adapts to context
  var a float32 = Tau          // ok
  var b float64 = Tau * 2      // ok, same constant

  const TauF float64 = 6.28318 // typed — now locked in
  var c float32 = TauF         // compile error: cannot use TauF (float64) as float32
  ```

- **No implicit conversions, ever.** Other languages promote `int → float` silently. Go does not. You write `float64(x) + y` explicitly. The verbosity is the point: every conversion is a place a precision loss or sign flip could happen, and the language wants those visible.

  ```go
  var i int = 3
  var f float64 = 1.5
  _ = i + f            // compile error: mismatched types int and float64
  _ = float64(i) + f   // ok — conversion is visible
  ```

- **`byte` vs `uint8`, `rune` vs `int32` — semantic aliases, not new types.** They exist purely to communicate *intent* at the declaration site. `rune` means "I'm holding a Unicode code point"; `int32` means "I'm holding a 32-bit signed integer that happens to fit in that range." The compiler doesn't care; the reader does.

  ```go
  var checksum uint8     // arithmetic on an 8-bit number
  var firstByte byte     // raw byte from a buffer
  var ch rune = '語'     // a Unicode code point
  // All three are uint8/uint8/int32 under the hood.
  ```

- **`iota` is a *line counter* inside a `const` block.** It's 0 on the first line, 1 on the second, and so on. It resets at the start of each `const (...)` group. The trick `1 << iota` gives you doubling values for bit flags. The pattern of leaving subsequent lines blank (just `B`, `C`) means "repeat the previous expression with the new `iota` value."

  ```go
  const (
      LogDebug = iota   // 0
      LogInfo           // 1 — expression "iota" repeats
      LogWarn           // 2
      LogError          // 3
  )

  const (
      PermRead    = 1 << iota   // 1
      PermWrite                 // 2
      PermExecute               // 4
  )
  ```

- **Unused *local variable* = compile error; unused *constant* = silently dropped.** A constant has no runtime cost — the compiler just doesn't emit it. A variable might have side effects (the RHS could be a function call), so leaving one unread is treated as a bug.

  ```go
  func demo() {
      const unused = 42   // fine — vanishes at compile time
      x := compute()      // compile error if x is never read — compute() might mutate state
  }
  ```

- **Strings are immutable byte sequences. `[]byte(s)` and `[]rune(s)` are different conversions.** Indexing a string gives you bytes, not characters. This lesson only scratches the surface — L03/L04 go deeper.

  ```go
  s := "héllo"
  len(s)          // 6 — byte count, not character count ('é' is 2 bytes)
  s[1]            // 0xA9 — a byte, not 'é'
  []rune(s)[1]    // 'é' — decoded code point
  ```

## What's in this directory

- `tempconv.go` — package `main`. Skeleton for a temperature converter that defines constants for scale identifiers via `iota`, a few well-known temperatures as typed/untyped constants, and conversion functions between Celsius, Fahrenheit, and Kelvin. Most of it is `// TODO`.

## Exercise

You're building a tiny temperature converter that forces you to confront `var` vs `:=` vs `const`, typed vs untyped constants, and explicit conversion.

1. In `tempconv.go`, declare an enumeration of temperature scales using a `const (...)` block and `iota`. Name them `Celsius`, `Fahrenheit`, `Kelvin` (these are the *scale identifiers*, not temperatures). Give the underlying constant a type — call it `Scale` and define it as `type Scale int`. Why give it a type? See gotcha #2.

2. Declare three well-known temperatures as **untyped** constants at package level: `AbsoluteZeroC = -273.15`, `FreezingC = 0`, `BoilingC = 100`. Do not annotate types. You'll see why in step 4.

3. Write three functions: `CtoF(c float64) float64`, `CtoK(c float64) float64`, `FtoC(f float64) float64`. Implement the math. Inside the functions, use `:=` for any intermediate locals.

4. Write a `func describe(s Scale, value float64) string` that returns a string like `"100.00 °C"` for `(Celsius, 100)`. Use a `switch` on `s` to pick the suffix. The `°` character is a non-ASCII rune — note that Go source is UTF-8, so writing `"°C"` in source is fine; you don't need an escape.

5. In `main`, use a mix of declaration forms deliberately:
   - One `var x float64` with no initializer (to demonstrate zero value).
   - One `x := someExpression()` (short declaration with inferred type).
   - One `var y int32 = 42` followed by an attempt to do arithmetic mixing `y` and an `int` literal — observe the error, then fix it with an explicit conversion.
   - Print `CtoF(BoilingC)` and `CtoK(AbsoluteZeroC)` using `describe`.

6. **Reflection (write your answers as comments in `tempconv.go`):**
   - Line 1: in step 5, when you tried to mix `y` (an `int32`) with an `int` literal — what was the *exact* compiler error? Why does the same arithmetic work fine with the untyped constants from step 2?
   - Line 2: when you used `iota` in step 1, which value does `Kelvin` end up with, and why? (Don't guess — work it out from the rules.)

7. **Stretch:** Add a fourth constant `Rankine` to the `iota` block. Now make `BoilingC` *typed*: `const BoilingC float32 = 100`. Try to pass it to `CtoF` which expects `float64`. Read the error. Then revert to untyped and observe that the same call now compiles. What did typing the constant cost you?

## Gotchas to watch for

- **"I'll just use `:=` everywhere."** You can't at package level — it's a function-scoped operator. And inside functions, `:=` will *silently shadow* an outer variable if you're not careful (we'll hit shadowing in L04). When in doubt — especially for zero-value declarations — `var` is clearer.

  ```go
  package foo

  count := 0   // compile error: syntax error, non-declaration statement outside function body
  var count = 0  // ok
  ```

- **"Why give the `iota` constants a named type (`type Scale int`)?"** Without it, `Celsius`, `Fahrenheit`, etc. are just `int`s, and the compiler will let you pass *any* `int` to `describe`. With a named type, the compiler enforces that only the named constants (or explicitly converted ints) can be passed. This is Go's main tool for enum-like safety.

  ```go
  type Direction int
  const ( North Direction = iota; East; South; West )

  func move(d Direction) { /* ... */ }

  move(North)    // ok
  move(2)        // compile error: cannot use 2 (int) as Direction
  move(Direction(2))  // ok — but you had to say it out loud
  ```

- **"Typed constant feels safer."** Often it isn't — it's just more restrictive. `const Pi float32 = 3.14159` cannot be added to a `float64` without a conversion. The untyped `const Pi = 3.14159` adapts to context. Default to untyped; type only when you want to *prevent* a specific category of use.
- **"`int` and `int32` are the same thing on a 64-bit machine."** No. `int` is a *distinct* type whose size happens to match the platform's word size. The compiler will refuse `var x int32 = 1; var y int = x` without an explicit conversion, even on a 32-bit platform where they have identical layout. Treat them as unrelated types that happen to share a representation.

  ```go
  var a int32 = 1
  var b int = a              // compile error: cannot use a (int32) as int
  var b int = int(a)         // ok
  ```

- **Floats are not exact.** `0.1 + 0.2 != 0.3` in Go too. Don't use `==` to compare floats; if you need precision (money!), reach for a decimal library — not built in.

  ```go
  fmt.Println(0.1 + 0.2)              // 0.30000000000000004
  math.Abs(0.1+0.2-0.3) < 1e-9        // true — compare with a tolerance
  ```

- **`rune` is `int32`, but a string's length in `len(s)` is its byte count, not rune count.** `len("é")` is 2, not 1. You'll feel this in L03; for now, just be aware that indexing a string gives bytes. (See `aside-bytes-and-unicode.md` in this directory for the deep version.)

## Checkpoint

You're done when:
- `go run ./ch02` prints both conversions correctly formatted (e.g. `212.00 °F`, `-459.67 °F` or whichever you computed).
- You can explain, without looking, the difference between `const x = 10` and `const x int = 10` and give an example where each is the right choice.
- You can articulate why `var y int32 = 42; var z = y + 1` works but `var y int32 = 42; var z int = y + 1` does not.
- You've written the two reflection comments at the top of `tempconv.go` and they're accurate.

Then run `/go:complete 02`.
