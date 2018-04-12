package ramda

// Equals returns true if its arguments are equivalent, false otherwise. Handles
// cyclical data structures.
func Equals(a ...interface{}) func() interface{} {
	if len(a) == 1 {
		return func() interface{} {
			return Head(a)
		}
	}

	return func() interface{} {
		for i := 1; i < len(a); i++ {
			if a[i] != Head(a) {
				return false
			}
		}
		return true
	}
}
