package trie

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRadixTree_Insert(t *testing.T) {
	tree := NewRadixTree()
	tree.Insert("/user/list", 1)
	tree.Insert("/user/create", 2)
	tree.Insert("/index", 3)
	assert.Equal(t, 1, tree.GetValue("/user/list"))
	assert.Equal(t, 2, tree.GetValue("/user/create"))
	assert.Equal(t, 3, tree.GetValue("/index"))
}
