package ramda

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHead(t *testing.T) {
	t.Run("[]bool", func(t *testing.T) {
		assert.Equal(t, true, Head([]bool{true, false}))
	})
	t.Run("[]int", func(t *testing.T) {
		assert.Equal(t, 1, Head([]int{1, 2}))
	})
	t.Run("[]string", func(t *testing.T) {
		assert.Equal(t, "a", Head([]string{"a", "b"}))
	})
	t.Run("[]int8", func(t *testing.T) {
		assert.Equal(t, int8(1), Head([]int8{1, 2}))
	})
	t.Run("[]int16", func(t *testing.T) {
		assert.Equal(t, int16(1), Head([]int16{1, 2}))
	})
	t.Run("[]int32", func(t *testing.T) {
		assert.Equal(t, int32(1), Head([]int32{1, 2}))
	})
	t.Run("[]int64", func(t *testing.T) {
		assert.Equal(t, int64(1), Head([]int64{1, 2}))
	})
	t.Run("[]uint", func(t *testing.T) {
		assert.Equal(t, uint(1), Head([]uint{1, 2}))
	})
	t.Run("[]uint8", func(t *testing.T) {
		assert.Equal(t, uint8(1), Head([]uint8{1, 2}))
	})
	t.Run("[]uint16", func(t *testing.T) {
		assert.Equal(t, uint16(1), Head([]uint16{1, 2}))
	})
	t.Run("[]uint32", func(t *testing.T) {
		assert.Equal(t, uint32(1), Head([]uint32{1, 2}))
	})
	t.Run("[]uint64", func(t *testing.T) {
		assert.Equal(t, uint64(1), Head([]uint64{1, 2}))
	})
	t.Run("[]uintptr", func(t *testing.T) {
		assert.Equal(t, uintptr(1), Head([]uintptr{1, 2}))
	})
	t.Run("[]float32", func(t *testing.T) {
		assert.Equal(t, float32(1), Head([]float32{1, 2}))
	})
	t.Run("[]float64", func(t *testing.T) {
		assert.Equal(t, float64(1), Head([]float64{1, 2}))
	})
	t.Run("[]complex64", func(t *testing.T) {
		assert.Equal(t, complex64(1), Head([]complex64{1, 2}))
	})
	t.Run("[]complex128", func(t *testing.T) {
		assert.Equal(t, complex128(1), Head([]complex128{1, 2}))
	})
	t.Run("string", func(t *testing.T) {
		assert.Equal(t, "1", Head("123"))
	})
	t.Run("[]interface{}", func(t *testing.T) {
		fn1 := Always(true)
		fn2 := Always(false)
		assert.Equal(t, fn1(), Head([]interface{}{fn1(), fn2()}))
	})
	t.Run("no match", func(t *testing.T) {
		assert.Equal(t, nil, Head([]http.Request{http.Request{}, http.Request{}}))
	})
}

func TestTail(t *testing.T) {
	t.Run("[]bool", func(t *testing.T) {
		assert.Equal(t, false, Tail([]bool{true, false}))
	})
	t.Run("[]int", func(t *testing.T) {
		assert.Equal(t, 2, Tail([]int{1, 2}))
	})
	t.Run("[]string", func(t *testing.T) {
		assert.Equal(t, "b", Tail([]string{"a", "b"}))
	})
	t.Run("[]int8", func(t *testing.T) {
		assert.Equal(t, int8(2), Tail([]int8{1, 2}))
	})
	t.Run("[]int16", func(t *testing.T) {
		assert.Equal(t, int16(2), Tail([]int16{1, 2}))
	})
	t.Run("[]int32", func(t *testing.T) {
		assert.Equal(t, int32(2), Tail([]int32{1, 2}))
	})
	t.Run("[]int64", func(t *testing.T) {
		assert.Equal(t, int64(2), Tail([]int64{1, 2}))
	})
	t.Run("[]uint", func(t *testing.T) {
		assert.Equal(t, uint(2), Tail([]uint{1, 2}))
	})
	t.Run("[]uint8", func(t *testing.T) {
		assert.Equal(t, uint8(2), Tail([]uint8{1, 2}))
	})
	t.Run("[]uint16", func(t *testing.T) {
		assert.Equal(t, uint16(2), Tail([]uint16{1, 2}))
	})
	t.Run("[]uint32", func(t *testing.T) {
		assert.Equal(t, uint32(2), Tail([]uint32{1, 2}))
	})
	t.Run("[]uint64", func(t *testing.T) {
		assert.Equal(t, uint64(2), Tail([]uint64{1, 2}))
	})
	t.Run("[]uintptr", func(t *testing.T) {
		assert.Equal(t, uintptr(2), Tail([]uintptr{1, 2}))
	})
	t.Run("[]float32", func(t *testing.T) {
		assert.Equal(t, float32(2), Tail([]float32{1, 2}))
	})
	t.Run("[]float64", func(t *testing.T) {
		assert.Equal(t, float64(2), Tail([]float64{1, 2}))
	})
	t.Run("[]complex64", func(t *testing.T) {
		assert.Equal(t, complex64(2), Tail([]complex64{1, 2}))
	})
	t.Run("[]complex128", func(t *testing.T) {
		assert.Equal(t, complex128(2), Tail([]complex128{1, 2}))
	})
	t.Run("[]interface{}", func(t *testing.T) {
		fn1 := Always(true)
		fn2 := Always(false)
		assert.Equal(t, fn2(), Tail([]interface{}{fn1(), fn2()}))
	})
	t.Run("no match", func(t *testing.T) {
		assert.Equal(t, nil, Tail([]http.Request{http.Request{}, http.Request{}}))
	})
}

func TestIndexOf(t *testing.T) {
	t.Run("[]bool", func(t *testing.T) {
		assert.Equal(t, 0, IndexOf(true, []bool{true, false}))
	})
	t.Run("[]int", func(t *testing.T) {
		assert.Equal(t, 0, IndexOf(1, []int{1, 2}))
	})
	t.Run("[]string", func(t *testing.T) {
		assert.Equal(t, 0, IndexOf("a", []string{"a", "b"}))
	})
	t.Run("[]int8", func(t *testing.T) {
		assert.Equal(t, 1, IndexOf(int8(2), []int8{1, 2}))
	})
	t.Run("[]int16", func(t *testing.T) {
		assert.Equal(t, 1, IndexOf(int16(2), []int16{1, 2}))
	})
	t.Run("[]int32", func(t *testing.T) {
		assert.Equal(t, 1, IndexOf(int32(2), []int32{1, 2}))
	})
	t.Run("[]int64", func(t *testing.T) {
		assert.Equal(t, 0, IndexOf(int64(1), []int64{1, 2}))
	})
	t.Run("[]uint", func(t *testing.T) {
		assert.Equal(t, 0, IndexOf(uint(1), []uint{1, 2}))
	})
	t.Run("[]uint8", func(t *testing.T) {
		assert.Equal(t, 0, IndexOf(uint8(1), []uint8{1, 2}))
	})
	t.Run("[]uint16", func(t *testing.T) {
		assert.Equal(t, 0, IndexOf(uint16(1), []uint16{1, 2}))
	})
	t.Run("[]uint32", func(t *testing.T) {
		assert.Equal(t, 0, IndexOf(uint32(1), []uint32{1, 2}))
	})
	t.Run("[]uint64", func(t *testing.T) {
		assert.Equal(t, 0, IndexOf(uint64(1), []uint64{1, 2}))
	})
	t.Run("[]uintptr", func(t *testing.T) {
		assert.Equal(t, 0, IndexOf(uintptr(1), []uintptr{1, 2}))
	})
	t.Run("[]float32", func(t *testing.T) {
		assert.Equal(t, 0, IndexOf(float32(1), []float32{1, 2}))
	})
	t.Run("[]float64", func(t *testing.T) {
		assert.Equal(t, 0, IndexOf(float64(1), []float64{1, 2}))
	})
	t.Run("[]complex64", func(t *testing.T) {
		assert.Equal(t, 0, IndexOf(complex64(1), []complex64{1, 2}))
	})
	t.Run("[]complex128", func(t *testing.T) {
		assert.Equal(t, 0, IndexOf(complex128(1), []complex128{1, 2}))
	})
	t.Run("no match", func(t *testing.T) {
		assert.Equal(t, -1, IndexOf(complex128(4), []complex128{1, 2}))
	})
}
