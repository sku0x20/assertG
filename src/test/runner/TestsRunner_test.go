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
	r.Add(T1)
	r.Run()
	if t1 == nil {
		t.Fatalf("didn't ran the tests")
	}
}

func Test_Fixtures(t *testing.T) {
	r := runner.NewTestsRunner[any](t)
	r.Setup(Setup)
	r.Add(T1)
	r.Teardown(Teardown)
	r.Run()
	if setupT == nil || teardownT == nil {
		t.Fatalf("didn't ran the fixtures")
	}
}

func Test_FixturesTMatches(t *testing.T) {
	r := runner.NewTestsRunner[any](t)
	r.Setup(Setup)
	r.Add(T1)
	r.Teardown(Teardown)
	r.Run()
	if setupT != t1 || teardownT != t1 {
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

var setupT *testing.T = nil

func Setup(t *testing.T) any {
	setupT = t
	return nil
}

var t1 *testing.T = nil

func T1(t *testing.T, extra any) {
	t1 = t
}

var teardownT *testing.T = nil

func Teardown(t *testing.T) {
	teardownT = t
}
