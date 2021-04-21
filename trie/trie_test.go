package trie

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrie_Insert(t *testing.T) {
	trie := New()
	trie.Insert("Hello")
}

func TestTrie_Search(t *testing.T) {
	trie := New()
	words := []string{"sex","sleep","son"}
	for _, word := range words {
		trie.Insert(word)
	}

	for _, word := range words {
		assert.True(t, trie.Search(word))
	}

	assert.False(t, trie.Search("se"))
}