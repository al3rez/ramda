package ramda

// Always returns a function that always returns the given value.
func Always(a interface{}) func() interface{} {
	return func() interface{} {
		return a
	}
}

// Identity a function that does nothing but return the parameter supplied to
// it. Good as a default or placeholder function.
func Identity(a interface{}) interface{} {
	return a
}
