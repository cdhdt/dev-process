// Package selfcheck is a throwaway Go module that exists only to exercise
// .github/workflows/go-ci.yml from inside this repository.
//
// It is not part of any parent Go module: the go.mod at the repository root
// belongs to this package alone (dev-process has no other Go code), so
// nothing here can leak into, or be affected by, a project that merely calls
// go-ci.yml with `uses:`. A consumer such as cdhdt/lapigo never checks out
// this repository's tree, so this module's go.mod has no effect on any
// consumer's own `go build ./...`.
//
// It deliberately does not live under testdata/: the go tool excludes any
// directory literally named "testdata" from a "./..." wildcard match (see
// `go help packages`), and go-ci.yml's steps all use "./...". A package
// placed there would compile locally and be invisible to CI, which is worse
// than not testing it at all.
//
// The package is deliberately tiny: it is the load-bearing wire in a smoke
// test, not a demonstration of Go. See plain_test.go and postgres_test.go
// for the two paths it exercises.
package selfcheck
