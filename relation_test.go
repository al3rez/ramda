package ramda

import (
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
