package sort

// QuickSort 快速排序
func QuickSort(items []int)   {
	quickSort(items, 0, len(items)-1)
}

// quickSort 快速排序
func quickSort(items []int, left, right int)  {
	low := left // 临时变量,低位下标
	high := right // 临时变量，高位下标
	if low >  high { // 作为截止条件
		return
	}
	flag := items[low] // 设置基数，用于对比
	for low < high { // 循环条件
		for low < high && items[high] >= flag { // 从右往左遍历，直到发现比基数小的数
			high--
		}

		items[low] = items[high] // 将较小的数放在低位地址
		for low < high && items[low] <= flag { // 从左往右遍历，直到发现比基数大的数
			low++
		}
		items[high] = items[low] // 将较大的树放在高位地址
	}
	items[low] = flag // 经过一轮遍历后，将基数赋值给低位地址，形成左边为比基数小的数组，右边为比基数大的数组
	quickSort(items, left, low-1) // 递归，分解
	quickSort(items, low+1, right) // 递归，分解
}
