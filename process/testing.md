# Testing

**A test that cannot fail on the bug it exists to catch is worse than no test,
because it is believed.** Everything below follows from that one sentence.

## Write the failing test first

Write it, watch it fail, then implement. A test written after the code passes
on its first run, which tells you nothing: you have not seen it able to fail.

For a bug fix this is not optional — reproduce the defect with a failing test
against the unfixed code before writing the fix. A fix without a reproduction
is a guess that happened to make a symptom go away.

## A guard does not exist until it has been seen to refuse

Reading a guard's code proves nothing. A green pipeline proves nothing. Build
the input that must be rejected, watch the rejection, and put the trace in the
pull request.

*Two ways this goes wrong, both observed.*

A generator emitted SQL that was asserted byte-for-byte against committed files
and never executed against a database. The pipeline had provisioned a database
server from its first run and nothing read the variable naming it. "It renders"
had been standing in for "it applies" for as long as the package existed. The
files did apply — established by hand, once, during review. Nothing would have
re-established it.

The fix introduced a second failure of the same shape: the new tier **skipped**
when the database variable was absent, and the pipeline runs tests without
`-v`, so a skip prints nothing. The day the variable stopped arriving, the tier
would vanish and the pipeline would stay green. A guard must fail where its
precondition was promised, and only skip where it genuinely cannot run.

## Assert on exact values, never on substrings

```go
if !strings.Contains(got, "file.yaml:2:1:") { }  // passes on almost anything
if reason == "" { }                              // says nothing about the message
if d.EndColumn != d.Pos.Column+1 { }             // passes when both are wrong
```

- Compare the **whole** output to a literal wherever the output is a contract:
  diagnostics, generated code, SQL, API responses.
- Never assert a field against another field of the same value. That passes
  when both are wrong together.
- An error message is part of the contract. Assert its exact text.
- Do not test the language. `x := S{F: f}` then `x.F == f` passes on the day
  the real invariant breaks.

*The measurement behind this.* A mutation pass over two packages: sixteen of
thirty-six introduced bugs survived. Replacing a caret-count clamp with a
million failed nothing. The handful of tests pinning a full expected string
caught nearly every mutant; the ones checking `Contains` caught almost none.

## Mutate the code to test the test

Apply the bug, confirm the suite goes red, revert. Do it for the branch you
just covered. This is the only direct measurement of whether a test can fail,
and it costs a minute.

Coverage is not this measurement. A line can be fully covered by a test that
would accept any result it produced.

## Invariants live in code, not in comments

A comment saying a slice "is sorted by name" with nothing sorting it, or "at
most one per entity" with nothing counting, is a wish.

If an invariant matters, something checks it: a constructor that is the only
way to build the value, or an explicit freeze step with a test that violates
the invariant and expects the failure.

Where a check cannot exist yet — the code it would read has not been written —
record the debt where the work that closes it will be planned, not only in a
comment beside the assertion. A comment is read by whoever is already looking
at that line, which is nobody until it is too late.

## Mechanical traps

**Cached results.** Test runners cache by content. A cached pass after a change
elsewhere is indistinguishable from a real one. Disable the cache in the
command the project's `check` target runs — in Go, `go test -count=1`.

**Local verification that skips a CI step.** The project's `check` target runs
what CI runs, in the same order. A local check missing a step is how a red
pipeline reaches the default branch with everyone believing it is green. When
CI gains a step, `check` gains it the same day.

**Conclusions drawn from a search that found nothing.** A negative result from
a search tool is only evidence if you have seen that tool find the thing.
Verify the pattern matches something you know is present before concluding it
matches nothing anywhere.
