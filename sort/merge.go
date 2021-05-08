// 归并排序
// 分解：将n个元素分成两个含n/2个元素的子序列
// 解决：用合并排序法对两个子序列进行递归排序
// 合并：合并两个已排序的子序列，得到最终结果
package sort

func MergeSort(items []int)  {
	mergeSort(items, 0, len(items))
}
// 自顶向下归并排序，排序范围在[left,right)的数组
func mergeSort(items []int, left, right int)  {
	// 元素数量大于1时才进入递归
	if right - left > 1 {
		// 将数组一分为二
		middle := (left + right) / 2

		// 先将左边数组进行排序
		mergeSort(items, left, middle)
		// 再将右边数组进行排序
		mergeSort(items, middle, right)
		// 合并
		merge(items, left, middle, right)
	}

}
// 归并
func merge(arr []int, low, middle, high int)  {
	leftLen := middle-low
	rightLen := high-middle
	// 使用临时数组来存储将要排序的元素
	temp := make([]int, 0, leftLen+rightLen)

	// 遍历对比两个数组的每个元素
	l, r := 0, 0
	for l < leftLen && r < rightLen{
		lValue := arr[low+l] // 左边数组的元素
		rValue := arr[middle+r] // 右边数组的元素
		if lValue < rValue {
			temp = append(temp, lValue)
			l++
		}else {
			temp = append(temp, rValue)
			r++
		}
	}

	// 将剩下的元素追加到临时数组里
	temp = append(temp, arr[low + l:middle]...)
	temp = append(temp, arr[middle+r:high]...)

	// 将已排序的数据复制到原数组里
	for idx, val := range temp {
		arr[low+idx] = val
	}
}
