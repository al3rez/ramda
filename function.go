package ramda

// Always returns a function that always returns the given value. Note that for
// non-primitives the value returned is a reference to the original value.
func Always(a interface{}) func() interface{} {
	return func() interface{} {
		return a
	}
}
