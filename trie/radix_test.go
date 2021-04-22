package trie

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"sync"
	"testing"
)

func TestRadixTree_ComparePrefix(t *testing.T) {
	trie := NewRadixTree()

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

func BenchmarkRadixTree_Insert(b *testing.B) {
	trie := NewRadixTree()
	for i := 0; i < b.N; i++ {
		trie.Insert("hello" + strconv.Itoa(i))
	}
}

func TestRadixTree_Insert(t *testing.T) {
	trie := NewRadixTree()
	words := []string{"sex","sleep","son", "se"}
	for _, word := range words {
		trie.Insert(word)
	}
	fmt.Println(trie)
}

func TestRadixTree_Search(t *testing.T) {
	trie := NewRadixTree()
	words := []string{"s","sex","sleep","son", "se"}
	for _, word := range words {
		trie.Insert(word)
	}
	for _, word := range words {
		assert.True(t, trie.Search(word))
	}

	assert.False(t, trie.Search("hello"))
}

func BenchmarkRadixTree_Search(b *testing.B) {
	trie := NewRadixTree()
	words := []string{"s","sex","sleep","son", "se"}
	for _, word := range words {
		trie.Insert(word)
	}
	for i := 0; i < b.N; i++ {
		trie.Search("sleep")
	}
}


func TestRadixTree_Delete(t *testing.T) {
	trie := NewRadixTree()
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