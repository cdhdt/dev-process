# Git conventions

Form only: how a branch is named, how a commit reads, how a pull request
closes. The cycle these sit inside is in [`workflow.md`](workflow.md).

## Branches

```
type/short-subject-N
```

`type` is one of `feat`, `fix`, `chore`, `ci`, `docs`, `refactor`, `test`.
`short-subject` is a few hyphenated words. `N` is the issue number, so the
branch names its ticket without opening the tracker.

Example: `fix/empty-target-relation-41`.

## Commits

Conventional prefix, imperative mood — what the commit does, not what it did.

```
type: imperative message
```

- Good: `fix: reject a belongsTo whose target is the empty string`
- Bad: `fix: fixed the empty target bug` (not imperative)

The body carries the reasoning. A commit that changes behaviour explains what
was wrong, not only what is now different — the diff already shows the latter.

## Identity

Project commits use a GitHub `noreply` address. Set it **per repository**, not
only globally:

```bash
git config user.email <id>+<user>@users.noreply.github.com
```

A global setting is one clone on one machine away from being wrong. A local one
travels with the repository you actually commit to.

**Squash merges do not use your git config.** GitHub composes them server-side
from your account's public email, so a repository whose every branch commit is
clean can still carry a personal address on every merge commit. Enable *Keep my
email address private* in account settings; a local `git config` cannot reach
this.

*This was found by auditing a repository the day before it went public.* Every
branch commit used the noreply address and every squash merge did not.

## Pull requests

One pull request per issue, closing it with `Closes #N` in the body. A pull
request touching two tickets is split in two.

The description says what changed and **why**. If it changes generated output
or user-visible behaviour, it shows a before/after sample — a reviewer should
not have to run the code to see what users will receive.

**Squash merge, always.** The default branch carries one commit per issue, not
one per intermediate step of the session that wrote it. The squash message
takes the pull request title.

**Delete the branch on merge.** A merged branch has no further use and only
lengthens the list.

## Protected default branch

Two conditions before a merge, no exceptions:

1. CI is green.
2. A review is approved, at the level the subject's criticality sets — see
   [`review.md`](review.md).

No `--no-verify`, no direct push to the protected branch, no merge without
review, outside the bootstrap commit documented in
[`workflow.md`](workflow.md).

**State this as enforced only when the platform enforces it.** Branch
protection is free on public repositories; on private ones it may require a
paid plan. Where the platform cannot enforce it, say so in the project's
`CONTRIBUTING.md` rather than letting a contributor believe a guard exists that
does not. A promised rule and an enforced rule are not the same thing, and the
difference matters exactly when someone is in a hurry.
