# Aside — Bytes, integers, Unicode, and UTF-8 (with Go in mind)

This is background reading for L02. The lesson touches `byte`, `rune`, `int` vs `int32`, and string indexing without explaining the foundations underneath. If those felt hand-wavy, read this; otherwise skip it.

## 1. Bits and bytes

A **bit** is one binary digit: 0 or 1. A **byte** is eight bits glued together, so it can hold any of 2⁸ = 256 distinct values. That's the universal unit of storage — disk, memory, network — every modern system addresses memory one byte at a time.

A byte is just a pattern of 8 bits. It has no inherent meaning. The same byte `01000001` could be:

- the integer 65,
- the ASCII character `'A'`,
- one byte of a larger multi-byte number,
- one byte of a UTF-8 sequence,
- a pixel value, a flag set, an opcode...

Meaning comes from *how the program interprets it*. The byte itself is dumb.

In Go: the type `byte` is an alias for `uint8` — they're literally the same type, and you pick the name based on what you're communicating. `byte` says "raw storage / part of a string / I/O." `uint8` says "I'm doing 8-bit unsigned arithmetic."

## 2. Signed vs. unsigned integers

Given 8 bits, you have 256 patterns. You can map those 256 patterns to numbers in two ways:

**Unsigned (`uint8`):** all 256 patterns represent non-negative integers — `0` to `255`. The pattern `00000000` is 0, `11111111` is 255. Straightforward.

**Signed (`int8`):** the 256 patterns are split: half for non-negative, half for negative. Range: `-128` to `127`. Go (like nearly every modern language) uses **two's complement** encoding for this. The top bit acts as a sign indicator, but the encoding is more clever than "1 bit for sign, 7 for magnitude" — two's complement makes addition and subtraction work with the same hardware circuit for signed and unsigned numbers. You don't need to know the encoding to use Go, but two facts matter:

1. The most negative value has no positive counterpart. `int8` goes to `-128` but only to `+127`. `-(-128)` overflows.
2. Overflow wraps silently in Go. `var x int8 = 127; x++` gives you `-128`, not an error. The CPU doesn't trap; Go doesn't add a check.

### The Go integer zoo

| Type | Bits | Range |
|---|---|---|
| `int8`  / `uint8`  | 8  | `-128..127` / `0..255` |
| `int16` / `uint16` | 16 | `-32_768..32_767` / `0..65_535` |
| `int32` / `uint32` | 32 | `≈ ±2.1 billion` / `0..≈4.3 billion` |
| `int64` / `uint64` | 64 | `≈ ±9.2 × 10¹⁸` / `0..≈1.8 × 10¹⁹` |
| `int`   / `uint`   | platform-sized (32 or 64) | matches CPU word size |
| `uintptr` | platform-sized | big enough to hold a pointer's bits |

Plus two aliases that carry meaning rather than introduce new types:

- `byte` = `uint8` — "I'm raw byte data."
- `rune` = `int32` — "I'm a Unicode code point."

**Why so many?** Because the right size depends on what the number means. A count of items in memory? `int`. A network protocol field that's spec'd to be 16 bits? `uint16`. A byte of a file? `byte`. A character? `rune`. Go forces you to pick, and forces explicit conversions between sizes, because mixing them silently is how bugs hide.

`int` is the default for "just give me a counter." Use sized types when something outside the program (a file format, a wire protocol, hardware) dictates the width, or when you need negative numbers ruled out (`uint` for sizes you know can't go below zero — though Go's convention is to use `int` even for sizes, because `len()` returns `int`).

## 3. From bytes to characters: ASCII

Once you have bytes, the next question is: how do we represent text?

The original answer, from the 1960s, was **ASCII**: a 7-bit mapping of 128 characters to integers. `'A'` is 65, `'a'` is 97, `' '` is 32, `'0'` is 48, newline is 10. Since 7 bits fits in a byte (with one bit to spare), ASCII text could be stored one character per byte and exchanged between any two ASCII-aware systems.

This worked beautifully for English. For everyone else, it was a disaster. French, German, Russian, Chinese, Arabic, Hebrew, Hindi — none fit. So through the '80s and '90s, dozens of incompatible 8-bit extensions appeared (Latin-1, Windows-1252, Shift-JIS, GB2312, ...), each filling the upper 128 byte values with a different alphabet. A file's bytes meant nothing without knowing which encoding to apply. Email and the web were a mess.

## 4. Unicode: the universal character table

**Unicode** is the modern fix. It's a single, global registry assigning every character — across every script, plus emoji, math symbols, currency signs, and arrows — a unique number called a **code point**, written `U+XXXX` in hexadecimal.

- `'A'` → U+0041 (65)
- `'é'` → U+00E9 (233)
- `'語'` → U+8A9E (35486)
- `'🎉'` → U+1F389 (127881)

Code points go up to U+10FFFF — about 1.1 million. Most are unused; the assigned ones currently number around 150,000.

Unicode is **just the table**. It says "the character 'é' is number 233." It does *not* say how to store that number in bytes. That's where encoding comes in.

The smallest standard integer type that fits any code point is 32 bits — which is why `rune` in Go is `int32`. A rune holds a code point. Mechanically it's a number; semantically it's a character.

## 5. UTF-8: how Unicode actually lives in bytes

If every character is a number up to ~1 million, the naive encoding is "use 4 bytes per character." That's called UTF-32. It works, but it bloats English text 4× and breaks every existing tool that expects ASCII bytes.

**UTF-8** is the brilliant compromise:

- Code points `0..127` (all of ASCII) → encoded as **one byte**, identical to ASCII.
- Code points `128..2047` → **two bytes**.
- Code points `2048..65535` → **three bytes**.
- Code points `65536..1114111` → **four bytes**.

So existing ASCII files are *already valid UTF-8*. English text takes one byte per character. European text with accents takes 1-2 bytes per character. CJK takes 3. Emoji and rare scripts take 4. The total cost is small and the backwards compatibility is total.

UTF-8 is also **self-synchronizing**: the leading bits of each byte tell you whether it's a single-byte character, the start of a multi-byte sequence, or a continuation byte. If you drop into the middle of a UTF-8 stream, you can find the next character boundary by scanning forward a few bytes.

Go's design choice: **strings are UTF-8 byte sequences**. The source code `.go` files are UTF-8. String literals are UTF-8. `os.ReadFile` returns UTF-8 bytes (assuming the file is UTF-8). The whole ecosystem assumes it.

## 6. Putting it together in Go

Here's the layer cake with Go's types attached:

```
Layer            Go type      What it is
─────            ───────      ──────────
Storage bit      —            one binary digit
Storage byte     byte/uint8   8 bits; raw memory/disk/network
Encoded text     string       a sequence of bytes (UTF-8 by convention)
Character        rune/int32   one Unicode code point
Glyph            —            what you see on screen
                              (sometimes 1 rune, sometimes many)
```

Three Go idioms follow from this:

**`len(s)` is byte count, not character count.**

```go
s := "héllo"
len(s)               // 6 — five characters, but 'é' is 2 UTF-8 bytes
```

**Indexing a string gives bytes.**

```go
s[1]                 // a byte (uint8) — the first byte of 'é', not 'é' itself
```

**`range` over a string decodes UTF-8 into runes.**

```go
for i, r := range s {
    // i = byte offset where this character started
    // r = rune (code point) of the character
}
```

Want a slice of all the characters as runes? `[]rune(s)`. Want the bytes? `[]byte(s)`. Both are explicit conversions — Go won't do either silently.

**Want the character count?** `utf8.RuneCountInString(s)` from `unicode/utf8`. Not `len`.

## 7. Why this matters for L02

L02 introduces `byte` and `rune` as "semantic aliases" without unpacking what they're aliases *for*. The short version:

- `byte` (= `uint8`) means "I'm a unit of raw storage. Probably part of a string's UTF-8 encoding, or a chunk of I/O, or a byte of a protocol field."
- `rune` (= `int32`) means "I'm one Unicode code point. I represent a character."

The compiler doesn't distinguish them from their underlying types. The reader does. Picking the right alias is documentation that the compiler can't enforce but other humans will rely on.

## 8. Further reading

- Rob Pike's blog post **"Strings, bytes, runes and characters in Go"** (go.dev/blog/strings) — the canonical explanation, from the language's co-designer.
- Joel Spolsky, **"The Absolute Minimum Every Software Developer Absolutely, Positively Must Know About Unicode and Character Sets"** — language-agnostic but foundational.
- Donovan & Kernighan, *The Go Programming Language*, §3.5 ("Strings") — Go-specific, spec-level precision.
- Bodner, *Learning Go* (2nd ed.), Chapter 3, "Strings, Runes, and Bytes" section — modern idiom framing.
