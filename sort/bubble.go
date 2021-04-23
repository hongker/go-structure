// 实现冒泡排序
package sort

// Swap 交换
func Swap(arr []int, x, y int)  {
	arr[x] ^= arr[y]
	arr[y] ^= arr[x]
	arr[x] ^= arr[y]
}

// BubbleSort 冒泡排序,时间复杂度O(n^2)
// 如果是数组{n,n-1...,2,1},则遍历次数=n*(n-1)/2
func BubbleSort(arr []int)  {
	length := len(arr)
	if length <= 1 {
		return
	}

	for i := 0; i < length-1; i++ {
		for j := 0; j < length -i-1; j++ {
			if arr[j] > arr[j+1] {
				Swap(arr, j, j+1)
			}
		}
	}
}
