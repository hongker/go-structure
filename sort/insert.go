package sort

// InsertSort 插入排序，时间复杂度O(n^2)
func InsertSort(arr []int)  {
	length := len(arr)
	if length <= 1 {
		return
	}

	// 从第二个数开始遍历
	for i := 1; i < length; i++ {
		preIndex := i-1 // 记录前面的下标
		current := arr[i] // 记录当前值
		for preIndex>=0 && arr[preIndex] > current { // 与前一个值对比，如果更小，则前移
			arr[preIndex+1] = arr[preIndex] // 前置后移
			preIndex--
		}
		arr[preIndex+1] = current
	}
}
