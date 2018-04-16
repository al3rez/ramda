package ramda

// Head returns the first element of the given list or string. In some libraries
// this function is named first.
func Head(s interface{}) interface{} {
	switch v := s.(type) {
	case []bool:
		return v[0]
	case []string:
		return v[0]
	case []int:
		return v[0]
	case []int8:
		return v[0]
	case []int16:
		return v[0]
	case []int32:
		return v[0]
	case []int64:
		return v[0]
	case []uint:
		return v[0]
	case []uint8:
		return v[0]
	case []uint16:
		return v[0]
	case []uint32:
		return v[0]
	case []uint64:
		return v[0]
	case []uintptr:
		return v[0]
	case []float32:
		return v[0]
	case []float64:
		return v[0]
	case []complex64:
		return v[0]
	case []complex128:
		return v[0]
	case []interface{}:
		return v[0]
	default:
		return nil
	}
}

// Tail returns all but the first element of the given list or string (or object
// with a tail method).
func Tail(s interface{}) interface{} {
	switch v := s.(type) {
	case []bool:
		return v[len(v)-1]
	case []string:
		return v[len(v)-1]
	case []int:
		return v[len(v)-1]
	case []int8:
		return v[len(v)-1]
	case []int16:
		return v[len(v)-1]
	case []int32:
		return v[len(v)-1]
	case []int64:
		return v[len(v)-1]
	case []uint:
		return v[len(v)-1]
	case []uint8:
		return v[len(v)-1]
	case []uint16:
		return v[len(v)-1]
	case []uint32:
		return v[len(v)-1]
	case []uint64:
		return v[len(v)-1]
	case []uintptr:
		return v[len(v)-1]
	case []float32:
		return v[len(v)-1]
	case []float64:
		return v[len(v)-1]
	case []complex64:
		return v[len(v)-1]
	case []complex128:
		return v[len(v)-1]
	case []interface{}:
		return v[len(v)-1]

	default:
		return nil
	}
}

// IndexOf returns the position of the first occurrence of an item in an array,
// or -1 if the item is not included in the array. Equals is used to determine
// equality.
func IndexOf(a, s interface{}) int {
	switch v := s.(type) {
	case []bool:
		for i, e := range v {
			if Equals(a.(bool), e)().(bool) {
				return i
			}
		}
	case []string:
		for i, e := range v {
			if Equals(a.(string), e)().(bool) {
				return i
			}
		}
	case []int:
		for i, e := range v {
			if Equals(a.(int), e)().(bool) {
				return i
			}
		}
	case []int8:
		for i, e := range v {
			if Equals(a.(int8), e)().(bool) {
				return i
			}
		}
	case []int16:
		for i, e := range v {
			if Equals(a.(int16), e)().(bool) {
				return i
			}
		}
	case []int32:
		for i, e := range v {
			if Equals(a.(int32), e)().(bool) {
				return i
			}
		}
	case []int64:
		for i, e := range v {
			if Equals(a.(int64), e)().(bool) {
				return i
			}
		}
	case []uint:
		for i, e := range v {
			if Equals(a.(uint), e)().(bool) {
				return i
			}
		}
	case []uint8:
		for i, e := range v {
			if Equals(a.(uint8), e)().(bool) {
				return i
			}
		}
	case []uint16:
		for i, e := range v {
			if Equals(a.(uint16), e)().(bool) {
				return i
			}
		}
	case []uint32:
		for i, e := range v {
			if Equals(a.(uint32), e)().(bool) {
				return i
			}
		}
	case []uint64:
		for i, e := range v {
			if Equals(a.(uint64), e)().(bool) {
				return i
			}
		}
	case []uintptr:
		for i, e := range v {
			if Equals(a.(uintptr), e)().(bool) {
				return i
			}
		}
	case []float32:
		for i, e := range v {
			if Equals(a.(float32), e)().(bool) {
				return i
			}
		}
	case []float64:
		for i, e := range v {
			if Equals(a.(float64), e)().(bool) {
				return i
			}
		}
	case []complex64:
		for i, e := range v {
			if Equals(a.(complex64), e)().(bool) {
				return i
			}
		}
	case []complex128:
		for i, e := range v {
			if Equals(a.(complex128), e)().(bool) {
				return i
			}
		}
	}
	return -1
}
