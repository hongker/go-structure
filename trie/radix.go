package trie

import (
	"strings"
	"sync"
)

type RadixTree struct {
	mu sync.RWMutex
	children map[string]*RadixTree
	isWord   bool
}

// 查询共同前缀
func (rt *RadixTree) comparePrefix(s1, s2 string) int {
	max := len(s1)
	if l := len(s2); l < max {
		max = l
	}

	var i int
	for i = 0; i < max; i++ {
		if s1[i] != s2[i] {
			break
		}
	}

	return i
}

func (rt *RadixTree) Insert(word string) {
	rt.mu.Lock()
	defer rt.mu.Unlock()
	for s, child := range rt.children {
		idx := rt.comparePrefix(s, word)
		if idx == 0 {
			continue
		}

		samePrefix := word[0:idx]

		if len(s) == idx {
			if len(word) == idx {
				child.isWord = true
			} else {
				child.Insert(word[idx:])
				return
			}

		} else {
			// 删除原来的节点
			delete(rt.children, s)
			// 添加新的共同前缀节点
			node := new(RadixTree)
			node.children = make(map[string]*RadixTree, maxCap)
			if len(word) == idx {
				node.isWord = true
			} else {
				node.Insert(word[idx:])
			}
			node.Insert(s[idx:])
			rt.children[samePrefix] = node

			return
		}

	}

	node := new(RadixTree)
	node.children = make(map[string]*RadixTree, maxCap)
	node.isWord = true
	rt.children[word] = node

}

func NewRadixTree() *RadixTree {
	return &RadixTree{
		children: make(map[string]*RadixTree, maxCap),
		isWord:   false,
	}
}

func (rt *RadixTree) Search(word string) bool {
	rt.mu.RLock()
	defer rt.mu.RUnlock()
	for s, child := range rt.children {
		if s == word {
			return child.isWord
		}

		if strings.HasPrefix(word, s) {
			return child.Search(word[len(s):])
		}
	}

	return false
}

// Delete 删除
func (rt *RadixTree) Delete(word string) {
	rt.mu.Lock()
	defer rt.mu.Unlock()
	for s, child := range rt.children {
		if s == word {
			child.isWord = false
		}

		if strings.HasPrefix(word, s) {
			child.Delete(word[len(s):])
		}
	}
}
