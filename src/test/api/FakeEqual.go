package api

type FakeEqual struct {
	e bool
}

func (st *FakeEqual) Equal(_ any) bool {
	return st.e
}
