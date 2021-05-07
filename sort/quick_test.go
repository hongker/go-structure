package sort

import (
	"fmt"
	"testing"
)

func Test_quickSort(t *testing.T) {
	items := []int{5,1,3,9,8,2,6,4,7}
	QuickSort(items)
	fmt.Println(items)
}
