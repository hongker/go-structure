package sort

import "math"

// ShellSort 希尔排序，时间复杂度:O(n^2)
// 优先比较较远的元素，又叫缩小增量排序
func ShellSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}

	// 按照二分法来分割数组
	for gap := Floor(length/2); gap >0 ; gap = Floor(gap / 2){
		for i := gap; i < length; i++ {
			j := i
			current := arr[i]
			for j>= gap && current < arr[j-gap] {
				arr[j] = arr[j-gap]
				j = j - gap
			}
			arr[j] = current
		}
	}
}

func Floor(n int) int {
	return int(math.Floor(float64(n)))
}