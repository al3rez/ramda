package ramda

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEquals(t *testing.T) {
	t.Run("arguments are equal", func(t *testing.T) {
		assert.Equal(t, true, Equals(true, true)())
	})
	t.Run("arguments are not equal", func(t *testing.T) {
		assert.Equal(t, false, Equals(true, false)())
	})
	t.Run("single argument", func(t *testing.T) {
		assert.Equal(t, true, Equals(true)())
	})
}

func TestGt(t *testing.T) {
	t.Run("rune", func(t *testing.T) {
		assert.Equal(t, false, Gt('a', 'z'))
		assert.Equal(t, true, Gt('z', 'a'))
	})
	t.Run("int", func(t *testing.T) {
		assert.Equal(t, false, Gt(1, 2))
		assert.Equal(t, true, Gt(2, 1))
	})

	t.Run("int8", func(t *testing.T) {
		assert.Equal(t, false, Gt(int8(1), int8(2)))
		assert.Equal(t, true, Gt(int8(2), int8(1)))
	})
	t.Run("int16", func(t *testing.T) {
		assert.Equal(t, false, Gt(int16(1), int16(2)))
		assert.Equal(t, true, Gt(int16(2), int16(1)))
	})
	t.Run("int32", func(t *testing.T) {
		assert.Equal(t, false, Gt(int32(1), int32(2)))
		assert.Equal(t, true, Gt(int32(2), int32(1)))
	})
	t.Run("int32", func(t *testing.T) {
		assert.Equal(t, false, Gt(int32(1), int32(2)))
		assert.Equal(t, true, Gt(int32(2), int32(1)))
	})
	t.Run("int64", func(t *testing.T) {
		assert.Equal(t, false, Gt(int64(1), int64(2)))
		assert.Equal(t, true, Gt(int64(2), int64(1)))
	})
	t.Run("uint", func(t *testing.T) {
		assert.Equal(t, false, Gt(uint(1), uint(2)))
		assert.Equal(t, true, Gt(uint(2), uint(1)))
	})
	t.Run("uint8", func(t *testing.T) {
		assert.Equal(t, false, Gt(uint8(1), uint8(2)))
		assert.Equal(t, true, Gt(uint8(2), uint8(1)))
	})
	t.Run("uint16", func(t *testing.T) {
		assert.Equal(t, false, Gt(uint16(1), uint16(2)))
		assert.Equal(t, true, Gt(uint16(2), uint16(1)))
	})
	t.Run("no match", func(t *testing.T) {
		assert.Equal(t, false, Gt(http.Request{}, http.Request{}))
	})
}
