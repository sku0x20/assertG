package api

func NewFakeEqual(e bool) *FakeEqual {
	return &FakeEqual{e: e}
}

type FakeEqual struct {
	e bool
}

func (st *FakeEqual) Equal(_ any) bool {
	return st.e
}
