package ramda

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInc(t *testing.T) {
	t.Run("rune", func(t *testing.T) {
		assert.Equal(t, 'b', Inc('a'))
	})
	t.Run("int", func(t *testing.T) {
		assert.Equal(t, 2, Inc(1))
	})
	t.Run("int8", func(t *testing.T) {
		assert.Equal(t, int8(2), Inc(int8(1)))
	})
	t.Run("int16", func(t *testing.T) {
		assert.Equal(t, int16(2), Inc(int16(1)))
	})
	t.Run("int32", func(t *testing.T) {
		assert.Equal(t, int32(2), Inc(int32(1)))
	})
	t.Run("int64", func(t *testing.T) {
		assert.Equal(t, int64(2), Inc(int64(1)))
	})
	t.Run("uint", func(t *testing.T) {
		assert.Equal(t, uint(2), Inc(uint(1)))
	})
	t.Run("uint8", func(t *testing.T) {
		assert.Equal(t, uint8(2), Inc(uint8(1)))
	})
	t.Run("uint16", func(t *testing.T) {
		assert.Equal(t, uint16(2), Inc(uint16(1)))
	})
	t.Run("uint64", func(t *testing.T) {
		assert.Equal(t, uint64(2), Inc(uint64(1)))
	})
	t.Run("uintptr", func(t *testing.T) {
		assert.Equal(t, uintptr(2), Inc(uintptr(1)))
	})
	t.Run("float32", func(t *testing.T) {
		assert.Equal(t, float32(2), Inc(float32(1)))
	})
	t.Run("float64", func(t *testing.T) {
		assert.Equal(t, float64(2), Inc(float64(1)))
	})
	t.Run("no match", func(t *testing.T) {
		assert.Equal(t, 0, Inc(http.Request{}))
	})
}
