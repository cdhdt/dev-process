# The development cycle

Every task follows the same path, whoever walks it — a person, or one of the
agents that write much of this code. The cycle exists because work changes
hands: between sessions, between agents, between the author and a reviewer who
was not there. A cycle that only its author understands is not a cycle.

```
issue
  → worktree           (a dedicated working directory)
  → branch             (process/git.md)
  → code and tests     (process/testing.md)
  → journal kept current
  → draft pull request (opened at the first commit, not the last)
  → adversarial review by someone other than the author
  → squash merge, branch deleted
  → issue closed automatically
```

Each arrow is required. A pull request that skipped the worktree, or a session
that hands back without a journal entry, is unfinished work even when the code
is correct.

The last step depends on the first: the pull request body carries `Closes #N`,
which only works if the work started from an issue. No branch without a ticket.

## One worktree per task

Before the first line of code, create a working directory dedicated to the
task:

```bash
git fetch origin
git worktree add ../<repo>-<short-subject>-<N> -b <type>/<short-subject>-<N> origin/develop
cd ../<repo>-<short-subject>-<N>
```

**Why not a plain `git checkout`.** A clone has one working directory. Two
sessions sharing it fight over the same files: one switches branch and the
other's work vanishes from under it. Cloning afresh per task avoids that but
pays a full network fetch and a duplicated `.git` every time. A worktree shares
one object store while giving each task its own directory and branch.

*This has already happened here.* Two agents worked the same clone on one pull
request; nothing broke, and it was the second agent that noticed. Nothing broke
is not a control — it is the outcome you get until you don't.

**Corollary: the shared clone's state tells you nothing.** It is probably on
someone else's branch. Read from the remote — `git show origin/develop:<file>`
— not from the shared working directory.

**Cleanup is part of the task.** After the merge:

```bash
git worktree remove ../<repo>-<short-subject>-<N>
git worktree prune
```

`git worktree remove` refuses to delete a directory holding uncommitted
changes. That refusal is a feature — treat it as a signal that work is about to
be lost, never as an obstacle to force past. Eight abandoned worktrees
accumulated across sessions on one project before anyone swept them; each one
had to be checked for unique work before removal, which took longer than
removing them at the time would have.

## The journal

Every task issue carries a **Journal** section. Anyone who touches the task
updates it **before handing back** — finished, half done, or abandoned. Four
questions: what the context is, what is done, what remains, what was decided.

This is not paperwork. Work is picked up by a different session, sometimes a
different agent, sometimes weeks later. Without a current journal, resuming
means reading the whole diff and reconstructing decisions already settled —
reliably longer than writing the journal would have taken.

*What it costs to skip.* On one project, four separate documents still declared
that no implementation existed while nearly nine thousand lines of tested code
sat merged in the default branch. Handing the project to a new agent required
auditing every status claim in the repository first. The code was never the
problem; the record of it was.

**A session that ends without a current journal is an unfinished session**,
however good the code it leaves.

## Draft pull requests

The pull request opens at the **first** commit, in draft, not when the work is
done. That is what makes work in progress visible: another contributor sees
what already exists instead of starting over. Without it, two people take the
same ticket — which is how this rule was learned.

Mark it ready for review when CI is green.

## Escalate after three round trips

When a review sends a task back, the reviewer moves it. At the third round
trip, stop and raise the difficulty rather than trying a fourth time. Three
consecutive failures is not an execution problem; it is evidence the subject is
harder than it was estimated to be, and it needs a different level of attention
— see [`review.md`](review.md).

## Size and difficulty are different axes

Estimate size — files touched, round trips expected — separately from how
critical the subject is. Size does not predict difficulty. Creating the same
workflow in five repositories is large and purely mechanical. A twenty-line
shell guard can reopen five times through five different doors.

What makes a subject critical, and what follows from it, is in
[`review.md`](review.md).

## Bootstrap: the one exception

The cycle assumes a base branch exists. The very first commit of an empty
repository has nothing to branch from, so it goes straight to the default
branch. From the second commit onward the cycle applies with no further
exception.
