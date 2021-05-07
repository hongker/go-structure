package search

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	items := []int{1,2,3,4,5,6,7,8,9}
	idx := Search(items, 5)
	fmt.Println(idx)
}