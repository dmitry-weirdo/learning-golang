package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	var s Stack[int]
	assert.Equal(t, true, s.IsEmpty())

	s.Push(10)
	assert.Equal(t, false, s.IsEmpty())

	v, ok := s.Pop()
	assert.Equal(t, true, ok)
	assert.Equal(t, 10, v)
	assert.Equal(t, true, s.IsEmpty())

	v, ok = s.Pop()
	assert.Equal(t, false, ok)
	assert.Equal(t, 0, v)
	assert.Equal(t, true, s.IsEmpty())

	s.Push(1)
	s.Push(2)

	v, ok = s.Pop()
	assert.Equal(t, true, ok)
	assert.Equal(t, 2, v)
	assert.Equal(t, false, s.IsEmpty())

	v, ok = s.Pop()
	assert.Equal(t, true, ok)
	assert.Equal(t, 1, v)
	assert.Equal(t, true, s.IsEmpty())

	v, ok = s.Pop()
	assert.Equal(t, false, ok)
	assert.Equal(t, 0, v)
	assert.Equal(t, true, s.IsEmpty())
}
