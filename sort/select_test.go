package sort

import (
	"fmt"
	"testing"
)

func TestSelectSort(t *testing.T) {
	arr := []int{1,7,6,2,8,2,8,4}
	SelectSort(arr)
	fmt.Println(arr)
}
