package trie

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"sync"
	"testing"
)

func TestTrie_Insert(t *testing.T) {
	trie := New()

	wg := sync.WaitGroup{}
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func(i int) {
			trie.Insert("hello" + strconv.Itoa(i))
			wg.Done()
		}(i)

	}

	wg.Wait()

}

func BenchmarkTrie_Insert(b *testing.B) {
	trie := New()
	for i := 0; i < b.N; i++ {
		trie.Insert("hello" + strconv.Itoa(i))
	}
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