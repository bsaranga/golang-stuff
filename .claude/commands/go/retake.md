---
description: Archive a prior lesson completion and reframe the lesson for another pass
argument-hint: <lesson-number> (e.g. 01, 02)
---

Set up a retake of lesson `$ARGUMENTS`. This command does **not** quiz the learner — it archives the prior result and reshapes the lesson so the learner can come at it fresh. After this runs, the learner reworks the exercise and runs `/complete $ARGUMENTS` again.

Steps:

1. **Verify a prior completion exists.** Require `ch$ARGUMENTS/lesson_completion.md`. If it doesn't exist, stop and tell the learner to run `/complete $ARGUMENTS` first — there's nothing to retake.

2. **Read all prior iterations.** Read `ch$ARGUMENTS/lesson_completion.md` plus any `ch$ARGUMENTS/lesson_completion_*.md` from earlier retakes. Note across them:
   - Which questions / concepts the learner consistently scored low on (2/4 or below).
   - Which gotchas or mental-model points seem to keep tripping them.
   - Which areas they already nailed — don't reteach those, just confirm them in passing.

3. **Archive the current completion.** Determine the next index `K`: one more than the highest existing `lesson_completion_<n>.md` (so first retake → `_1`, second → `_2`, etc.). Rename `ch$ARGUMENTS/lesson_completion.md` → `ch$ARGUMENTS/lesson_completion_$K.md` with `git mv` (or plain `mv` if not tracked). Keep its `chmod 444` read-only status.

4. **Read the lesson as it stands.** Read `ch$ARGUMENTS/README.md` and the starter `.go` file(s). Read `lesson-plan.md`'s entry for L`$ARGUMENTS` for the canonical objective.

5. **Reframe the lesson — pedagogically, based on the weak areas from step 2.** The goal is a *different angle on the same objective*, not a harder or easier lesson. Pick whichever of these moves best targets the gaps; combine if it helps:
   - **Recast the exercise** around a different concrete scenario that forces the same idiom (e.g. if the lesson teaches slices via a number filter, retry it via a log-line parser). The mechanics must still exercise the same idiom the learner missed.
   - **Invert the framing.** If iteration 1 was "build X from scratch," iteration 2 might be "here is broken code that uses X — find and fix the bug." Bug-finding probes mental model differently than greenfield writing.
   - **Add a constraint** that makes the weak concept unavoidable (e.g. "this time, do it without a temp variable" / "do it without indexing into the slice").
   - **Sharpen the self-check questions** in the README to directly target the concepts the learner missed last time. Replace the weakest 1–2 base questions; keep the rest.
   - **Rewrite or expand the Gotchas / Key idioms sections** to spell out the specific misconception the prior `lesson_completion_*.md` files revealed — but don't hand them the answer; frame it as a question or a "watch for…" pointer.

   Do NOT change the lesson's core objective or stage. Do NOT solve the new exercise. Do NOT pre-answer the new self-check questions.

6. **Add a retake header to the README.** At the top of `ch$ARGUMENTS/README.md`, just under the lesson title, add or update a small block:

   ```markdown
   > **Retake iteration $K+1.** Prior attempts: `lesson_completion_1.md` … `lesson_completion_$K.md`.
   > Focus this pass: <1–2 sentences naming the specific concepts to nail this time, derived from prior weak areas>.
   ```

   If a previous retake header already exists, update it in place rather than stacking new ones.

7. **Reset starter code.** Rewrite the `.go` file(s) so they're back to a `// TODO`-marked skeleton that fits the new framing. The learner should be able to start from a blank-ish slate, not their previous solution. If their previous solution is committed in git, that's fine — they can diff later; do not preserve it in the working tree.

8. **Report back to the learner**:
   - The archived file path (`lesson_completion_$K.md`).
   - A short bulleted summary of the weak areas you identified across prior iterations.
   - What changed in the lesson (exercise reframing, new constraints, sharpened questions, etc.) and *why* — tie each change to a specific weak area.
   - A one-line pointer: "When you're done, run `/complete $ARGUMENTS` again."

Rules:
- Be honest about what was weak. Don't soften it — the point of a retake is to target the gap.
- Don't escalate difficulty for its own sake. The objective stays constant; only the angle changes.
- Don't commit. The learner reviews and commits when they're ready.
- Don't touch other lessons.
