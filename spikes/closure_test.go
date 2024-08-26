package spikes

import "testing"

func Test_defer_closure(tm *testing.T) {
	f := func(t *testing.T) {
		defer func() { t.Logf("defered closure") }()
		t.Logf("inside closure")
	}
	defer func() {
		tm.Logf("defered")
	}()
	tm.Run("inside-test", func(t *testing.T) {
		f(t)
		t.Logf("inside-test")
		defer func() {
			t.Logf("defered inside run")
		}()
	})
}
