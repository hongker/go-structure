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
	words := []string{"sex","sleep","son", "se"}
	for _, word := range words {
		trie.Insert(word)
	}

	for _, word := range words {
		assert.True(t, trie.Search(word))
	}

	assert.False(t, trie.Search("so"))
}

func TestTrie_Delete(t *testing.T) {
	trie := New()
	words := []string{"sex","sleep","son", "se"}
	for _, word := range words {
		trie.Insert(word)
	}
	assert.True(t, trie.Search("sex"))
	assert.True(t, trie.Search("se"))
	trie.Delete("sex")
	assert.False(t, trie.Search("sex"))
	assert.True(t, trie.Search("se"))
	trie.Delete("se")
	assert.False(t, trie.Search("se"))
}