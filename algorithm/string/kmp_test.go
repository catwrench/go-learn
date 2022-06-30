package string

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKmpBase(t *testing.T) {
	kmp2 := NewKmp("XXYXXYXXY")
	index2 := kmp2.KmpSearch("XXYXXYXXX", 0)
	assert.Equal(t, -1, index2)
}

func TestKmp(t *testing.T) {
	kmp := NewKmp("Hello world !")

	index := kmp.KmpSearch("wor", 0)
	assert.Equal(t, 6, index)
	index = kmp.KmpSearch("ll", 0)
	assert.Equal(t, 2, index)
	index = kmp.KmpSearch("lasd", 0)
	assert.Equal(t, -1, index)
	index = kmp.KmpSearch("Hello", 0)
	assert.Equal(t, 0, index)
	index = kmp.KmpSearch("Hello", 3)
	assert.Equal(t, -1, index)
	return
}

func TestKmp2(t *testing.T) {
	kmp := NewKmp("Hello world !")

	index := kmp.KmpSearch2("wor", 0)
	assert.Equal(t, 6, index)
	index = kmp.KmpSearch2("ll", 0)
	assert.Equal(t, 2, index)
	index = kmp.KmpSearch2("lasd", 0)
	assert.Equal(t, -1, index)
	index = kmp.KmpSearch2("Hello", 0)
	assert.Equal(t, 0, index)
	index = kmp.KmpSearch2("Hello", 3)
	assert.Equal(t, -1, index)
	return
}
