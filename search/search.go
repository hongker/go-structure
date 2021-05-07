package search

// Search 二分查找
// items为一个有序数组
func Search(items []int, num int) int {
	l := len(items)
	if items[0] > num || items[l-1] < num{
		return -1
	}

	if items[0] == num {
		return 0
	}

	if items[l-1] == num {
		return l-1
	}

	var left = 0
	var right = l-1
	for left < right {
		middle := (right - left) / 2
		if items[middle] == num {
			return middle
		}else if items[middle] > num {
			right = middle
		}else {
			left = middle+1
		}
	}
	return -1



}
