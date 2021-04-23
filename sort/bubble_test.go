package sort

import (
	"fmt"
	"testing"
)

func TestSwap(t *testing.T)  {
	arr := []int{1,2}
	Swap(arr, 0, 1)
	fmt.Println(arr)
}

func BenchmarkSwap(b *testing.B) {
	arr := []int{1,2}
	for i := 0; i < b.N; i++ {
		Swap(arr, 0, 1)
	}
}

func TestBubbleSort(t *testing.T) {

	arr := []int{}
	for i := 100; i >0 ; i-- {
		arr = append(arr, i)
	}
	BubbleSort(arr)
	fmt.Println(arr)
}