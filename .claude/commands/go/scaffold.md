---
description: Scaffold a new lesson directory (chNN/) from the lesson plan
argument-hint: <lesson-number> (e.g. 02, 03)
---

Scaffold lesson `$ARGUMENTS` for this Go curriculum.

Steps:
1. Read `lesson-plan.md` and find the entry for L`$ARGUMENTS`. Use its title and one-line description as the basis for the new lesson.
2. Read the relevant chapter(s) from the two reference PDFs in `references/` (Donovan & Kernighan — *The Go Programming Language*; Jon Bodner — *Learning Go*, 2nd ed.) to ground the lesson in authoritative material. Use `Read` with a targeted `pages:` range — do NOT load entire books. Prefer Bodner for idiom/style framing and Donovan for spec-level precision. Distill — never copy prose verbatim.
3. Read `ch01/README.md` to match the exact structure and tone (Objective → "You should finish this lesson able to answer..." numbered list of 4–6 self-check questions → Key idioms / mental model → What's in this directory → Exercise (numbered steps, optional **Stretch:**) → Gotchas to watch for → Checkpoint).
4. Create `ch$ARGUMENTS/README.md` following that template, tailored to the lesson's topic. Self-check questions should probe the *why*, not trivia.
5. Create starter `.go` file(s) under `ch$ARGUMENTS/` with `// TODO` markers where the learner fills in code. Keep the starter minimal and runnable as a skeleton (or clearly marked as non-compiling until completed).
6. Do NOT solve the exercise. Do NOT answer the self-check questions.
7. Do NOT commit — leave that to the learner.

After scaffolding, report: lesson title, files created, and a one-line pointer to where the learner should start.
