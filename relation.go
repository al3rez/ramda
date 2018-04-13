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

// Gt returns true if the first argument is greater than the second; false
// otherwise.
func Gt(a, b interface{}) bool {
	switch v := a.(type) {
	case int:
		return v > b.(int)
	case int8:
		return v > b.(int8)
	case int16:
		return v > b.(int16)
	case int32:
		return v > b.(int32)
	case int64:
		return v > b.(int64)
	case uint:
		return v > b.(uint)
	case uint8:
		return v > b.(uint8)
	case uint16:
		return v > b.(uint16)
	case uint64:
		return v > b.(uint64)
	case uintptr:
		return v > b.(uintptr)
	case float32:
		return v > b.(float32)
	case float64:
		return v > b.(float64)
	default:
		return false
	}
}
