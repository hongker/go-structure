package trie

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRadixTree_Insert(t *testing.T) {
	tree := NewRadixTree()
	tests := []struct {
		name string
		key string
		val interface{}
	}{
		{
			name: "number",
			key:  "ten",
			val:  10,
		},
		{
			name: "string",
			key:  "hello",
			val:  "world",
		},
		{
			name: "sex",
			key:  "sex",
			val:  "性别",
		},
	}
	for _, test := range tests {
		tree.Insert(test.key, test.val)
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.val, tree.GetValue(test.key))
		})
	}
}
