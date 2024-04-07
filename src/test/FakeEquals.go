package test

type FakeEquals struct {
	equal bool
}

func (st *FakeEquals) Equals(other any) bool {
	return st.equal
}
