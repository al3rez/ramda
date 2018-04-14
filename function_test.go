package ramda

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlways(t *testing.T) {
	t.Run("returns the given value", func(t *testing.T) {
		assert.Equal(t, true, Always(true)())
		assert.Equal(t, "", Always("")())
		assert.Equal(t, http.Request{}, Always(http.Request{})())
	})
}

func TestIdentity(t *testing.T) {
	t.Run("returns the given value", func(t *testing.T) {
		assert.Equal(t, true, Identity(true))
		assert.Equal(t, "", Identity(""))
		assert.Equal(t, http.Request{}, Identity(http.Request{}))
	})
}
