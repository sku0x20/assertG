package api

type FakeEquals struct {
	Equal bool
}

func (st *FakeEquals) Equals(other any) bool {
	return st.Equal
}

func (st *FakeEquals) String() string {
	if st.Equal {
		return "everything is equal"
	} else {
		return "nothing is equal"
	}
}
