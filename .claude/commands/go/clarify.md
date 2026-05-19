---
description: Scan lesson READMEs for learner-added questions and answer them inline
argument-hint: [lesson-number] (optional; defaults to all lessons)
---

Find and answer learner-added questions in lesson READMEs.

Scope:
- If `$ARGUMENTS` is given (e.g. `02`), only scan `ch$ARGUMENTS/README.md`.
- Otherwise scan every `chNN/README.md` in the repo.

For each README:
1. Identify the base self-check list (under "You should finish this lesson able to answer..."). **Leave it untouched.**
2. Look for learner-added questions — typically under a `## Questions` heading, or as extra items appended beyond the base list, or anywhere with an unanswered `?` that isn't part of the base self-check.
3. For each unanswered learner question, edit the README to add a tight, idiomatic answer directly beneath the question. Match the existing README's tone (concise, mental-model-first, no fluff, no restating the question).
4. Do NOT answer or modify the base self-check questions.
5. Do NOT solve the exercise for them — answer the conceptual question, don't write the code they're meant to write.

After editing, report: which files were modified and a one-line summary of the questions answered in each. If no learner questions were found, say so explicitly.
