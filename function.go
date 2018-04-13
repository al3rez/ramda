package ramda

// Always returns a function that always returns the given value.
func Always(a interface{}) func() interface{} {
	return func() interface{} {
		return a
	}
}
