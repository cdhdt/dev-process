package selfcheck

import "testing"

// TestAlwaysPasses is the entirety of the `postgres: false` path's payload:
// go-ci.yml's plain `check` job must be able to gofmt, vet, staticcheck,
// test, and build this package end to end with no database involved.
func TestAlwaysPasses(t *testing.T) {
	if 1+1 != 2 {
		t.Fatal("arithmetic broke")
	}
}
