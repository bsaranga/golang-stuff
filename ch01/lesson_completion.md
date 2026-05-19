# Lesson 01 — Completion

**Date:** 2026-05-19
**Final grade:** A (19/20)

## Q1. If you removed both aliases from the imports in `hello.go` (kept bare `import ".../ch01/messages"` and `import ".../ch01/notes"`), would the file compile? If not, what's the error and why?
**Answer:** It wouldn't compile, I'd get the error saying that `greet` has been redeclared and the Evening function is undefined. This is because `greet` is imported from messages and even though /notes share the same package name it should be aliased in order to be used, else a package conflict occurs.
**Rating:** 4/4 — identified the collision and the reason precisely.
**Correct take:** Both directories declare `package greet`, so two bare imports bind the same local identifier `greet` in this file → `greet redeclared in this block`. Exactly the situation aliases exist for.

## Q2. Of the three — directory name, `package X` declaration, import path — which two are coupled, which is independent, and which determines the call-site identifier?
**Answer:** Directory name and import path are coupled, each determines the other. The package name is independent from directory name and import path, and the package name is the identifier, which can also be aliased.
**Rating:** 4/4 — clean mapping of all three.
**Correct take:** Import path ↔ directory (coupled). `package X` is independent and is the default call-site identifier; an alias overrides it locally.

## Q3. What does capitalization-as-visibility give a reader at the call site that `public`/`private` keywords wouldn't?
**Answer:** The reader or a tool can immediately identify that it is an unexported method at the callsite, no need to go to the source and see whether it has been defined as private or public.
**Rating:** 3/4 — right mental model, slight slip (capitalized `Bar` in `foo.Bar` is *exported*, not unexported). Also didn't name the grep angle.
**Correct take:** Visibility lives at *every reference*, not just the declaration. `foo.Bar(x)` tells you at the call site that `Bar` is exported — no jump-to-definition. `grep -nE '[a-z]+\.[A-Z]'` finds every cross-package call in a tree.

## Q4. If two files in the same directory declared different packages (`package greet` and `package helpers`), what happens and why does Go enforce this?
**Answer:** The compiler would complain that it found multiple packages. This breaks the one package = one directory rule. When go resolves a package it uses the mod file module declaration and finds all modules relative to the root, all files in a directory is compiled to one package, so multiple package definitions doesn't fit this pattern.
**Rating:** 4/4 — correctly tied the rule to the compilation-unit invariant.
**Correct take:** Directory = compilation/linkage unit. Mixed `package X` declarations make the unit ambiguous — the toolchain wouldn't know what to link as. It's an invariant, not a style rule.

## Q5. If `ch01/notes/notes.go` were changed to `package notes` and the notes import were left unaliased (messages still aliased as `greet`), does it compile, and what do you write at the call site for Evening?
**Answer:** Yes, this is valid go and it will compile. I use notes.Evening(...) at the call site because it's a reference to the package. The package name is the identifier.
**Rating:** 4/4 — correct and reasoned from the right principle.
**Correct take:** Compiles. Call site is `notes.Evening(...)` — the unaliased import binds the local identifier to the declared package name, which is now `notes`. No collision with the aliased `greet`.

## Overall feedback
Mental model is locked in: you cleanly separated the four-way directory ↔ package ↔ import path ↔ call-site identifier chain, and you understood aliases as a caller-side disambiguator rather than a fix for mismatches. One small wobble on Q3 — re-read the "Capitalization as visibility, not just terseness" bullet in `ch01/README.md`: in `foo.Bar(x)`, the capital `B` marks `Bar` as *exported*, and the value is that you can see it without leaving the call site. Ready to move on.
