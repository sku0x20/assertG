package test

/*
func Test_WithoutRunner(t *testing.T) {
	c := assert.NewCaptureT(t)
	c.ThatAny(10).IsNotNil().IsEqualTo(10)
}
*/
/*
func Test_WithRunner(tm *testing.T) {
	r := runner.NewTestsRunner[any](tm)
	r.Setup(func(t *testing.T) any {
		return nil
	})
	r.Teardown(func(t *testing.T, extra any) {
		t.Logf("tear-down")
	})
	r.Add(func(t *testing.T, extra any) {
		assert.ThatAny(10).IsNotNil().IsEqualTo(10)
	})
	r.Run()
}
*/
