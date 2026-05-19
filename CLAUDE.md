# Working in this repo

This is a self-paced Go curriculum. `lesson-plan.md` defines L01–L53 across six stages. Each lesson lives in its own `chNN/` directory with a `README.md` (objective, key idioms, exercise, gotchas, checkpoint) and starter `.go` files containing `// TODO` markers.

The learner (Saranga) drives the pace and works each exercise themselves. Do not solve exercises pre-emptively.

## Lesson README question protocol

Every `chNN/README.md` contains a numbered list under "You should finish this lesson able to answer..." These are the learner's **self-check** questions. Leave them alone — the learner answers them on their own.

If something is unclear, the learner will **append additional questions** to that lesson's README (typically under a `## Questions` heading, or as extra numbered items beyond the base self-check list). Your job:

1. On any session touching this repo, scan every `chNN/README.md` for learner-added questions that are not yet answered.
2. Answer those questions **inline in the README itself** (edit the file), placing the answer directly under each question.
3. Keep answers tight, idiomatic, and matching the tone of the existing README. No fluff, no restating the question.
4. Never modify or answer the base self-check questions.

## Scaffolding new lessons

When asked to scaffold the next lesson (e.g. "scaffold L02"), create `chNN/` with:
- `README.md` following the L01 structure: Objective → "You should finish this lesson able to answer..." list → Key idioms / mental model → What's in this directory → Exercise → Gotchas → Checkpoint.
- Starter `.go` file(s) with `// TODO` markers where the learner fills in code.

## Commit cadence

Each completed lesson is committed before moving to the next, so progress is auditable in `git log`. Do not commit unless asked.
