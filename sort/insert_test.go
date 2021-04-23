package sort

import (
	"fmt"
	"testing"
)

func TestInsertSort(t *testing.T) {
	arr := []int{1,7,6,2,8,2,8,4}
	InsertSort(arr)
	fmt.Println(arr)
}