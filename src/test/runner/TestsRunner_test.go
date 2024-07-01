package runner

import (
	"assertG/src/pkg/runner"
	"testing"
)

func Test_Nop(t *testing.T) {
	r := runner.NewTestsRunner(t)
	r.Run()
}

func Test_NoFixtures(t *testing.T) {
	r := runner.NewTestsRunner(t)
	r.Add(ChangeRan)
	r.Run()
	if !ran {
		t.Fatalf("didn't ran the tests")
	}
}

func Test_Fixtures(t *testing.T) {
	r := runner.NewTestsRunner(t)
	r.Setup(ChangeSetup)
	r.Add(ChangeRan)
	r.Teardown(ChangeTeardown)
	r.Run()
	if !setup || !teardown {
		t.Fatalf("didn't ran the fixtures")
	}
}

var ran = false

func ChangeRan(t *testing.T) {
	ran = true
}

var setup = false

func ChangeSetup(t *testing.T) any {
	setup = true
	return nil
}

var teardown = false

func ChangeTeardown(t *testing.T) {
	teardown = true
}
