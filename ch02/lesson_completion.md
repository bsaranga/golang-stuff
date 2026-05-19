# Lesson 02 — Completion

**Date:** 2026-05-19
**Final grade:** B (16/20)

## Q1. Which of Go's four declaration forms is illegal at package level, and why is that restriction useful rather than annoying?
**Answer:** "x := v is restricted at package level, since package level does not have an order of statements like a function level does."
**Rating:** 4/4 — identifies the right form and the core reason (no execution order at package scope).
**Correct take:** `:=` is a *statement*, not a declaration. Package level has only declarations — a dependency-ordered set, not a sequence. Forcing `var`/`const` also makes the package's surface area visually obvious when scanning a file.

## Q2. Why does Go refuse `var x int = 1; var y float64 = x + 1.5` but accept `const c = 1; var y float64 = c + 1.5`?
**Answer:** "x is a fixed type (int) and it cannot be added to 1.5 resolving to a float64 — Go doesn't do implicit type conversions. An untyped constant implicitly takes the shape of types around it; c is added to 1.5 as a float64."
**Rating:** 4/4 — captures fixed-at-declaration vs deferred-until-use.
**Correct take:** `x`'s type is fixed the moment the `var` is parsed. An untyped constant has *no type* until used in a context that demands one; at that point it adopts whatever fits.

## Q3. What does `type Scale int` buy you over using plain `int` for the enum constants?
**Answer:** "Enum-like safety. If we call describe with any int we'd get a compiler error. If we drop the named type and use an int parameter, we lose enum safety and any int can be passed in."
**Rating:** 4/4 — correct on both directions.
**Correct take:** The named type makes `describe(s Scale, ...)` reject arbitrary `int`s — only the declared constants (or an explicit `Scale(n)` conversion) compile. Go's main tool for enum safety.

## Q4. Given `const ( A = 1 << iota; B; _; D )`, what are A, B, D?
**Answer:** "2, 4, 6, 8"
**Rating:** 1/4 — wrong operation (treated `1 << iota` as additive, not as a shift) and didn't notice `_` is a discard, not a name.
**Correct take:** `1 << 0 = 1`, `1 << 1 = 2`, `1 << 2 = 4` (discarded via `_`), `1 << 3 = 8`. So A=1, B=2, D=8. The expression repeats per line; `iota` provides the shift amount.

## Q5. If you'd written `var z = yInt + 1` instead of `var z int = int(yInt) + 1`, would it compile? What is z's type?
**Answer:** "Compiles, z would be int32 with value 43. When a variable is untyped, type conversion happens implicitly."
**Rating:** 3/4 — right answer, but the explanation calls `z` "untyped" — variables always have a type. The flexibility belongs to the untyped *constant* `1`.
**Correct take:** Compiles. `1` is an untyped constant that conforms to `yInt`'s type (`int32`), so `yInt + 1` is `int32`. `var z = ...` infers `z`'s type from the RHS, giving `z` type `int32`, value 43.

## Overall feedback
The declaration-form and untyped-constant mental models are solid — you can articulate the *when* of type fixing, which is the hard part of this lesson. The named-type enum reasoning is sharp. Two areas to revisit: (1) the `iota` bit-flag pattern — re-read the Key idioms snippet and convince yourself why `1 << iota` doubles (the shift amount is the line index); (2) the precise wording around "untyped" — that property belongs to *constants*, never to variables, because every variable has a type the moment it's declared. Variables can have their type *inferred*, but they're never untyped.
