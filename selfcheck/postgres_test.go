package selfcheck

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5"
)

// databaseURLEnv must match the `database-url-env` this repository's own
// self-test workflow (.github/workflows/selftest.yml) passes to the
// `postgres: true` call of go-ci.yml.
const databaseURLEnv = "SELFCHECK_TEST_DATABASE_URL"

// postgresJobID is the id go-ci.yml gives its own Postgres-backed job. It is
// not a private implementation detail: it is the second half of the check
// name ("ci / check-with-postgres") that cdhdt/lapigo's branch protection
// already depends on, so relying on it here is relying on a name go-ci.yml
// cannot rename without breaking a consumer regardless of this test.
const postgresJobID = "check-with-postgres"

// TestDatabaseURLReachesProcess is the reason this module exists. It proves
// that go-ci.yml's dynamic `$GITHUB_ENV` export -- the step that writes the
// caller's chosen variable name, because GitHub Actions does not allow an
// expression as an `env:` key -- actually reaches the test process under
// that caller-chosen name, in this case databaseURLEnv.
//
// This repository's own selftest.yml calls go-ci.yml twice from the same
// commit: once with `postgres: false`, once with `postgres: true`. Both
// calls build and test this exact package, so this test has to behave
// correctly under both without go-ci.yml's interface giving it a direct way
// to tell them apart -- there is no working-directory input, and the two
// calls check out identical content. GITHUB_JOB is what GitHub Actions sets
// to the id of the job whose steps are currently running; for a job
// dispatched through workflow_call that is go-ci.yml's own job id, "check"
// or "check-with-postgres", regardless of what this repository names the
// caller job. That is the one signal available at run time that
// distinguishes "this run promised a database" from "it never did" without
// touching go-ci.yml.
//
// Modeled on cdhdt/lapigo's internal/ddl/apply_test.go: a plain t.Skip on a
// missing variable prints nothing under `go test ./...` without -v, which is
// how CI runs, so a silently broken export would leave the pipeline green --
// the same failure shape this test exists to close. So: skip only where a
// database was genuinely never promised (postgres: false, or a developer
// running this module locally), fail where one was.
func TestDatabaseURLReachesProcess(t *testing.T) {
	if os.Getenv("GITHUB_JOB") != postgresJobID {
		t.Skip("not running under the " + postgresJobID + " job; no database was promised here")
	}

	url := os.Getenv(databaseURLEnv)
	if url == "" {
		t.Fatalf("%s is unset while running under %s: the dynamic $GITHUB_ENV export did not reach the test process", databaseURLEnv, postgresJobID)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	conn, err := pgx.Connect(ctx, url)
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	defer func() { _ = conn.Close(context.Background()) }()

	var one int
	if err := conn.QueryRow(ctx, "SELECT 1").Scan(&one); err != nil {
		t.Fatalf("query: %v", err)
	}
	if one != 1 {
		t.Fatalf("SELECT 1 returned %d, want 1", one)
	}
}
