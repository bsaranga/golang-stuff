---
description: Answer a learner question about a lesson, optionally enriching the README
argument-hint: <lesson-number> <question>
---

Answer a learner question in the context of a specific lesson.

Parse `$ARGUMENTS`:
- First token is the lesson number (e.g. `02` → `ch02/`).
- The remainder is the question.
- If the lesson number is missing or no question is supplied, ask the learner to provide them — don't guess.

Steps:

1. Read `chNN/README.md` (and skim the starter `.go` files in that directory if relevant) to ground yourself in what the lesson is actually teaching. Also consult `references/` (Bodner / Donovan) when the question touches spec-level semantics or idiom framing — use targeted `pages:` reads, not whole-book loads.

2. Decide where the answer should live:
   - **Chat-only** (default): if the question is a one-off clarification that doesn't materially improve the lesson for a future reader, answer directly in chat.
   - **Edit the README**: if the answer fills a genuine gap — a missing gotcha, an unclear idiom, a mental-model nuance the lesson should have surfaced — edit the README to fold the clarification in. Place it under the most relevant existing section (Gotchas, Key idioms, etc.), or add a new short subsection if nothing fits. Match the existing tone: concise, mental-model-first, no fluff, no restating the question.

   When in doubt, prefer chat-only. The README is a teaching artifact, not a Q&A log.

3. Never touch the base self-check list ("You should finish this lesson able to answer..."). Never solve the exercise for the learner — answer the conceptual question, don't write the code they're meant to write.

4. Report briefly: whether you answered in chat only or edited the README (and which section), in one line.
