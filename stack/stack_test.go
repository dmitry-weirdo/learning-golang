package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	var s Stack[int]

	s.Push(10)

	v, ok := s.Pop()
	assert.Equal(t, true, ok)
	assert.Equal(t, 10, v)

	v, ok = s.Pop()
	assert.Equal(t, false, ok)
	assert.Equal(t, 0, v)

	s.Push(1)
	s.Push(2)

	v, ok = s.Pop()
	assert.Equal(t, true, ok)
	assert.Equal(t, 2, v)

	v, ok = s.Pop()
	assert.Equal(t, true, ok)
	assert.Equal(t, 1, v)

	v, ok = s.Pop()
	assert.Equal(t, false, ok)
	assert.Equal(t, 0, v)
}
