---
description: Quiz the learner on a completed lesson and record a graded completion file
argument-hint: <lesson-number> (e.g. 01, 02)
---

Run an interactive completion check for lesson `$ARGUMENTS`.

Steps:

1. Read `ch$ARGUMENTS/README.md` and any `.go` files in `ch$ARGUMENTS/` so you understand what the learner just finished. If `ch$ARGUMENTS/lesson_completion.md` already exists, stop and tell the learner the lesson is already marked complete (don't overwrite).

2. Ask the learner **5 questions** about the lesson, one at a time, directly in the chat. Wait for an answer to each before asking the next — do not batch them. Mix the question types:
   - At least 2 should come from (or closely paraphrase) the base self-check list in the README.
   - At least 2 should probe the *why* / mental model behind a key idiom or gotcha from the lesson.
   - At least 1 should be a small applied question (e.g. "what would happen if you changed X to Y in the exercise code?").
   Do NOT show the answers up front. Do NOT give hints unless the learner explicitly asks.

3. After each answer, internally note a rating on this scale (do not reveal per-question scores until the end):
   - **4 — Solid**: correct, idiomatic, shows real understanding.
   - **3 — Mostly right**: correct core idea, minor gap or imprecise wording.
   - **2 — Partial**: touches the right area but misses the key point.
   - **1 — Off**: incorrect or fundamentally confused.
   - **0 — Skipped / "I don't know"**.

4. Once all 5 questions are answered, compute the final grade out of 20 and map to a letter:
   - 18–20 → **A** · 15–17 → **B** · 12–14 → **C** · 8–11 → **D** · <8 → **Needs review**

5. Write `ch$ARGUMENTS/lesson_completion.md` with this structure:

   ```markdown
   # Lesson $ARGUMENTS — Completion

   **Date:** <today's date>
   **Final grade:** <letter> (<score>/20)

   ## Q1. <question>
   **Answer:** <learner's answer verbatim>
   **Rating:** <n>/4 — <one-line justification>
   **Correct take:** <tight, idiomatic correct answer, 1–3 lines>

   ## Q2. ...
   (repeat for Q1–Q5)

   ## Overall feedback
   <2–4 sentences: what's solid, what to revisit. Reference specific README sections or files in chNN/ if relevant.>
   ```

6. Make the file read-only: `chmod 444 ch$ARGUMENTS/lesson_completion.md`.

7. Report the final grade and the path to the completion file. Do NOT commit.

Rules:
- Be honest with ratings. This is a self-check tool; inflating grades defeats the point.
- Keep "Correct take" tight and matched to the README's tone — no fluff, no restating the question.
- Never reveal the correct answer before the learner has answered.
- If the learner abandons mid-quiz, do not write the file.
