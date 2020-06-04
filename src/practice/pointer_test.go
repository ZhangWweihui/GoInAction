package practice

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPointer1(t *testing.T) {
	a := 20
	t.Logf("a value = %d, type = %T", a, a)

	p := &a
	t.Logf("p value = %x, type = %T", p, p)

	assert.Equal(t, a, *p)

	var ptr *int
	//assert.Equal(t, ptr, 0)
	//assert.Equal(t, ptr, nil)
	assert.True(t, ptr == nil)
}
