package _spikes

import "testing"

type S struct {
}

func newS() *S {
	return &S{}
}

func newSNil() *S {
	return nil
}

func Test_cast(t *testing.T) {
	s := newS()
	t.Logf("%v", s)
	nilS := newSNil()
	t.Logf("%v", nilS)
	if nilS == nil {
		t.Logf("nilS is nil")
	}
	var anyS any = newSNil()
	t.Logf("%v", anyS)
	if nilS == nil {
		t.Logf("anyS is nil")
	}
	var sSNil *S = newSNil()
	t.Logf("%v", sSNil)
	if nilS == nil {
		t.Logf("sSNil is nil")
	}
	var sNilAnyCast any = newSNil()
	casted := sNilAnyCast.(*S)
	t.Logf("%v", casted)
	if casted == nil {
		t.Logf("casted is nil")
	}
	var sNilAnyCast2 any = newSNil()
	casted2, ok := sNilAnyCast2.(*S)
	t.Logf("%v", casted2)
	t.Logf("%v", ok)
	if !ok {
		t.Logf("ok is false")
	}
	if casted2 == nil {
		t.Logf("casted2 is nil")
	}
}
