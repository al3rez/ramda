package ramda

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnd(t *testing.T) {
	t.Run("given true and true", func(t *testing.T) {
		assert.Equal(t, true, And(true, true))
	})
	t.Run("given true and false and vice versa", func(t *testing.T) {
		assert.Equal(t, false, And(true, false))
		assert.Equal(t, false, And(false, true))
	})
	t.Run("given false and false", func(t *testing.T) {
		assert.Equal(t, false, And(false, false))
	})
}

func TestCond(t *testing.T) {
	fn := NewCond(Cond{
		Equals(0), Always("water freezes at 0°C"),
		Equals(100), Always("water boils at 100°C"),
		T(), Always("nothing special happens at %d°C"),
	})

	t.Run("water temperature", func(t *testing.T) {
		t.Run("given 0°C", func(t *testing.T) {
			assert.Equal(t, "water freezes at 0°C", fn(0))
		})
		t.Run("given 50°C", func(t *testing.T) {
			assert.Equal(t, "nothing special happens at 50°C", fn(50))
		})
		t.Run("given 100°C", func(t *testing.T) {
			assert.Equal(t, "water boils at 100°C", fn(100))
		})
	})
}

func TestT(t *testing.T) {
	t.Run("returns true", func(t *testing.T) {
		assert.Equal(t, true, T()())
	})
}
