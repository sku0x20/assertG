package test

type FakeEquals struct {
	equal bool
}

func (st *FakeEquals) Equals(other any) bool {
	return st.equal
}

func (st *FakeEquals) String() string {
	if st.equal {
		return "everything is equal"
	} else {
		return "nothing is equal"
	}
}
