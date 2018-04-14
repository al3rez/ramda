package ramda

// Inc increments its argument.
func Inc(a interface{}) interface{} {
	switch v := a.(type) {
	case int:
		return v + 1
	case int8:
		return v + int8(1)
	case int16:
		return v + int16(1)
	case int32:
		return v + int32(1)
	case int64:
		return v + int64(1)
	case uint:
		return v + uint(1)
	case uint8:
		return v + uint8(1)
	case uint16:
		return v + uint16(1)
	case uint64:
		return v + uint64(1)
	case uintptr:
		return v + uintptr(1)
	case float32:
		return v + float32(1)
	case float64:
		return v + float64(1)
	default:
		return 0
	}
}
