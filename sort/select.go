package sort

// SelectSort 选择排序，时间复杂度O(n^2)
func SelectSort(arr []int)  {
	length := len(arr)
	if length <= 1 {
		return
	}
	// 循环遍历
	for i := 0; i < length - 1; i++ {
		// 以arr[i]为基数，找到比它小的数进行交换
		for j := i+1; j < length; j++ {
			if arr[i] > arr[j] {
				Swap(arr, i, j)
			}
		}
	}
}
