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

## Why the rules carry their incidents

Every rule here is followed by what happened when it was missing. That is not
storytelling. A rule stated bare reads as ceremony and gets skipped under
pressure; the same rule with the failure it prevents survives contact with a
deadline. Where an incident is cited, it is real and its evidence is public.

## Licence

MIT. See [`LICENSE`](LICENSE).
