---
description: Explain a Go term in context, cross-referenced with the reference books
argument-hint: <term> (e.g. "iota", "rune", "interface satisfaction")
---

Explain the Go term/concept `$ARGUMENTS` in the context of this curriculum.

Steps:

1. Identify what `$ARGUMENTS` refers to in Go. If it's ambiguous (e.g. "make" could mean the builtin or the Unix tool), pick the Go-language meaning and note the ambiguity in one line.

2. Look up the term in both reference PDFs in `references/`:
   - **Donovan & Kernighan — *The Go Programming Language*** (spec-level precision; memory model, semantics, edge cases).
   - **Jon Bodner — *Learning Go*, 2nd ed.** (idiom/style framing; modern usage, when-to-reach-for-it).
   Use `Read` with a targeted `pages:` range — do NOT load whole books. If you don't know which chapter to look in, first read the relevant table-of-contents pages (Bodner TOC: PDF pages ~9–13; Donovan TOC: near the front of its PDF) to find the right range, then jump there.

3. Write the explanation directly to chat (do NOT create a file). Structure:

   ```
   # `$ARGUMENTS`

   **One-line definition:** <tight, idiomatic, single sentence>

   **In context:** <2–4 sentences on what it actually does in Go, why the
   language has it, and the mental model — not a textbook recitation>

   **Gotchas / common confusions:** <bullet list, 2–4 items max, only
   include if there's something genuinely worth flagging>

   **Tiny example:** <a 3–8 line Go snippet that illustrates the term in
   use — runnable or near-runnable, no fluff>

   **References:**
   - Bodner, *Learning Go* 2nd ed., Ch X "<chapter name>", pp. <page range>
   - Donovan & Kernighan, *The Go Programming Language*, §X.Y "<section>", pp. <page range>
   ```

   Omit a section only if it genuinely doesn't apply (e.g. no meaningful gotchas). Cite both books when both cover the term; cite one if only one does. Use the book's printed page numbers, not the PDF page numbers.

4. Match the README tone elsewhere in this repo: concise, mental-model-first, no marketing fluff, no restating the question. Quote from the books only if a specific phrase is load-bearing — otherwise distill in your own words.

5. If `$ARGUMENTS` is not a Go concept (e.g. the user typed a typo or something unrelated), say so plainly and suggest the closest Go term you can match.

Do NOT modify any lesson files. Do NOT commit.
