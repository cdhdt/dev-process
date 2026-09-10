# dev-process

The development process my projects follow, written once so each repository
does not reinvent it.

This repository is public for one reason: **a rule that decides the fate of a
contribution has to be readable by the person contributing.** A project that
tells you to follow its process and then hides it is running a closed trial.

## What is here

| Document | What it settles |
|---|---|
| [`process/workflow.md`](process/workflow.md) | The cycle a task goes through, from issue to merge |
| [`process/git.md`](process/git.md) | Branch names, commit format, how a pull request closes |
| [`process/review.md`](process/review.md) | What gets reviewed how hard, and what a reviewer owes the author |
| [`process/testing.md`](process/testing.md) | The bar a test has to clear to count as one |

Plus reusable CI workflows under [`.github/workflows/`](.github/workflows),
callable from any repository with `uses:`.

## What is not here

The tooling that carries this process — agent definitions, model routing,
bootstrap scripts, machine-specific notes — is private. It says nothing a
contributor needs, and publishing it would expose infrastructure rather than
explain a rule.

The split is deliberate and runs one way: **anything here may be referenced
from a private repository; nothing here references a private one.** A public
document that links into a private repository hands the reader a 404, which is
the same failure as not publishing the process at all.

## Using it in a project

Each repository states the rules that bind a contributor in its own
`CONTRIBUTING.md` — a contributor should never have to follow an external link
to learn how to name a branch. This repository holds the long form: the
reasoning, the incidents that produced each rule, and the parts that are
identical across projects.

Where a project's own document disagrees with this one, the project wins. Its
`CONTRIBUTING.md` and `CLAUDE.md` are closer to the code and are what a
reviewer will actually apply.

## Versioning the reusable workflows

The workflows under [`.github/workflows/`](.github/workflows) are consumed by
`uses:` from other repositories' own CI. `cdhdt/lapigo` calls `go-ci.yml` and
its branch protection requires the resulting `check-with-postgres` job to be
green — a mistake in a workflow here reaches that repository's pull requests
before anyone here has looked at it.

**Pin `@v1`, never `@main`.**

```yaml
jobs:
  ci:
    uses: cdhdt/dev-process/.github/workflows/go-ci.yml@v1
```

`@main` moves on every merge to this repository, with no boundary between
"someone edited a workflow" and "every consumer's next CI run uses the new
one." `@main` is **not supported** for a consumer's CI: if a workflow file
here changes underneath you and nobody said so, it is because you pinned a
branch instead of a version.

### What `v1` promises

`v1` is a moving tag over the current major version. Re-pointing it to a new
commit on `main` must never require a consumer to change anything:

- The `workflow_call` `inputs:` of an existing workflow keep their names,
  types, and defaults. Adding an optional input with a default is fine;
  renaming, removing, or narrowing one is not.
- A job id a consumer's branch protection depends on (`check`,
  `check-with-postgres`) keeps its id and keeps meaning the same thing.
- The steps a workflow runs may be improved — a newer tool version, a faster
  cache, an extra check that was silently missing — as long as a green run
  before the change stays green after it, for any consumer already passing.

In short: **`v1` may get stricter or faster, never incompatible.**

### What would justify `v2`

A new major tag is for a change that a consumer must react to, for example:

- Removing, renaming, or repurposing a `workflow_call` input.
- Renaming a job id that branch protection rules reference.
- Changing what a passing run means in a way that could fail a consumer that
  was previously green through no fault of their own (for example, turning on
  a lint category that a real project is expected to violate).

A breaking change lands on `main` and ships under `v2`; `v1` keeps pointing at
the last commit that honoured the promise above. Both tags can coexist
indefinitely — a consumer upgrades to `v2` on their own schedule, by editing
one line.

### Keeping `v1` pointed at `main`

A tag that moves only when someone remembers to move it is worse than no tag:
a consumer silently freezes on whatever commit `v1` last pointed to, and
nothing here says so. [`update-major-tag.yml`](.github/workflows/update-major-tag.yml)
re-points `v1` at the tip of `main` automatically on every push to `main`,
using the repository's own `GITHUB_TOKEN`. No maintainer step, nothing to
forget. It also accepts a manual `workflow_dispatch` (with the tag name as an
input, defaulting to `v1`) so the mechanism can be exercised without a push to
`main` — this is how it was verified before this policy was documented; see
the pull request that introduced it.

## Why the rules carry their incidents

Every rule here is followed by what happened when it was missing. That is not
storytelling. A rule stated bare reads as ceremony and gets skipped under
pressure; the same rule with the failure it prevents survives contact with a
deadline. Where an incident is cited, it is real and its evidence is public.

## Licence

MIT. See [`LICENSE`](LICENSE).
