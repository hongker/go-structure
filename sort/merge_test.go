package sort

import (
	"fmt"
	"testing"
)

func TestMergeSort(t *testing.T) {
	items := []int{5,1,3,9,8,2,6,4,7,10}
	MergeSort(items)
	fmt.Println(items)
}