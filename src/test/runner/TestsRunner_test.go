package runner

import (
	"assertG/src/pkg/runner"
	"testing"
)

func Test_Nop(t *testing.T) {
	r := runner.NewTestsRunner[any](t)
	r.Run()
}

func Test_NoFixtures(t *testing.T) {
	r := runner.NewTestsRunner[any](t)
	r.Add(ChangeRan)
	r.Run()
	if testT == nil {
		t.Fatalf("didn't ran the tests")
	}
}

func Test_Fixtures(t *testing.T) {
	r := runner.NewTestsRunner[any](t)
	r.Setup(ChangeSetup)
	r.Add(ChangeRan)
	r.Teardown(ChangeTeardown)
	r.Run()
	if setupT == nil || teardownT == nil {
		t.Fatalf("didn't ran the fixtures")
	}
}

func Test_FixturesTMatches(t *testing.T) {
	r := runner.NewTestsRunner[any](t)
	r.Setup(ChangeSetup)
	r.Add(ChangeRan)
	r.Teardown(ChangeTeardown)
	r.Run()
	if setupT != testT || teardownT != testT {
		t.Fatalf("t different for fixtures")
	}
}

func Test_Extras(tm *testing.T) {
	r := runner.NewTestsRunner[string](tm)
	r.Setup(func(ts *testing.T) string {
		return "some value"
	})
	r.Add(func(tr *testing.T, extra string) {
		if extra != "some value" {
			tr.Fatalf("wrong value, expected \"some value\", got \"%s\"", extra)
		}
	})
	r.Run()
}

var testT *testing.T = nil

func ChangeRan(t *testing.T, extra any) {
	testT = t
}

var setupT *testing.T = nil

func ChangeSetup(t *testing.T) any {
	setupT = t
	return nil
}

var teardownT *testing.T = nil

func ChangeTeardown(t *testing.T) {
	teardownT = t
}
