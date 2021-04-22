package trie

import "sync"

type nodeType int
const(
	rootNode nodeType = iota // 根节点
	staticNode				 // 子节点
)

const(
	maxCap = 26
)

// Trie 字典树
type Trie struct {
	// lock for map
	mu sync.RWMutex
	// 子节点
	children map[rune]*Trie
	// 节点类型
	isWord bool
}

// Insert 插入元素
func (trie *Trie) Insert(word string) {
	trie.mu.Lock()
	defer trie.mu.Unlock()
	for _, w := range word {
		if trie.children[w] == nil { // 不存在子节点
			// 初始化新的子节点
			node := new(Trie)
			node.children = make(map[rune]*Trie, maxCap)
			trie.children[w] = node
		}
		// 存在子节点,继续遍历
		trie = trie.children[w]
	}
	// 遍历结束,设置结束标识
	trie.isWord = true
}
// Search 查询
func (trie *Trie) Search(word string) bool{
	trie.mu.RLock()
	defer trie.mu.RUnlock()
	// 遍历字符串
	for _, w := range word {
		// 如果发现有字符不存在，则说明不匹配
		if trie.children[w] == nil {
			return false
		}
		// 继续遍历
		trie = trie.children[w]
	}
	// 遍历完整个字符后，且同时满足字符串已结束，说明匹配成功
	// 如果存在sleep，而搜索的sle,则为false
	return trie.isWord
}

// Delete 删除
func (trie *Trie) Delete(word string)  {
	trie.mu.RLock()
	defer trie.mu.RUnlock()
	// 遍历字符串
	for _, w := range word {
		// 如果发现有字符不存在，则说明不匹配
		if trie.children[w] == nil {
			return
		}
		// 继续遍历
		trie = trie.children[w]
	}

	trie.isWord = false
}

// New 实例化
func New() *Trie {
	return &Trie{
		children:  make(map[rune]*Trie, maxCap),
		isWord: false,
	}
}