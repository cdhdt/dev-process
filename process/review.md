# Review

## Review is adversarial, and never by the author

A reviewer's job is to **refute** the work, not to bless it: to find the input
that breaks it, the case the tests do not reach, the claim in the description
that the code does not support. Approval is what is left when refutation
failed.

The reviewer is always someone other than whoever wrote the code. When agents
write the code, this means a different agent with no memory of writing it — an
author re-reading their own work re-runs the reasoning that produced the bug
and finds it sound, because it was sound to them an hour ago.

*What this catches.* On one project, an adversarial pass over a package that
looked finished — full statement coverage, a green pipeline — returned three
real defects. One emitted a foreign key column with no constraint and no
diagnostic: a schema that validated clean and produced an unconstrained
database. Coverage had not noticed, because the line ran; it just did the wrong
thing.

## What makes a subject critical

Critical regardless of size:

- authentication, sessions, identity linking
- personal data
- money
- irreversible migrations
- **any guard whose failure is silent**

Everything else is standard. The last entry is the one people forget, and it is
the one that bites: a guard that fails loudly gets fixed the same morning. A
guard that stops running keeps a pipeline green while protecting nothing, and
is discovered months later or not at all.

Criticality decides how hard a change is reviewed and by what. **How that maps
to reviewer capability is a project's own concern** and is not specified here —
it depends on tooling that changes faster than this document should.

Criticality is independent of size. A one-line change to a security guard is
critical; a thousand lines of mechanical transcription is not.

## What a reviewer verifies rather than reads

Reading code and observing a green pipeline are not verification. A review that
concludes from either has checked that the work looks right.

- **Re-run what the author claims.** A description saying the suite passes, the
  output is byte-identical, or a benchmark improved is a claim. Run it.
- **Make the guard fail.** See [`testing.md`](testing.md) — a guard nobody has
  watched refuse anything is a guard nobody has tested.
- **Check what the tests cannot fail on.** Mutate the code the test covers and
  confirm the test goes red. Coverage says a line ran, never that a wrong
  result would have been caught.

*Worth the reminder:* an adversarial pass on a well-tested package found that
sixteen of thirty-six deliberately introduced bugs survived its suite. Deleting
the handling that one function existed for failed nothing.

## What a reviewer owes the author

- **Findings, with a reproduction.** "This looks fragile" is not a finding.
  What input, what state, what wrong output.
- **A verdict.** Blocking or not. A list of observations that does not say
  whether the work can merge leaves the author to guess.
- **Correction when the review is wrong.** A reviewer who was mistaken says so
  plainly and moves on. Reviews are not scored.

An author receiving a review is not required to agree. Disagreement is settled
by evidence — reproduce it or refute it — not by seniority, and not by
implementing a suggestion the author believes is wrong to end the exchange.

## Three round trips

At the third return to the author, stop. Raise the subject's criticality and
change the level of attention rather than attempting a fourth pass. Three
consecutive failures measure the problem, not the person working it.
