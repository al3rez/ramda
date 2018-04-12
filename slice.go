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
