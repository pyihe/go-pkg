package sorts

import (
	"sort"
)

// Bubble 冒泡排序，每次找出最小的值放在最前面
// 平均时间复杂度：O(n^2)
// 最坏时间复杂度：O(n^2)
// 最优时间复杂度：O(n)
// 空间复杂度：O(1)
// 稳定性：YES
func Bubble(data sort.Interface) {
	// isChange 是否有发生了交换动作, 如果没有发生交换，证明数据已经有序了, 不需要在进行遍历比较
	isChange, n := false, data.Len()
	if n <= 1 {
		return
	}
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if data.Less(j, j+1) {
				data.Swap(j, j+1)
				isChange = true
			}
		}
		if !isChange {
			break
		}
	}
}

// Select 选择排序:每次选择出未排序切片里最大或者最小的数放入已排好序的数组里
// 平均时间复杂度：O(n^2)
// 最坏时间复杂度：O(n^2)
// 最优时间复杂度：O(n^2)
// 空间复杂度: O(1)
// 稳定性: NO
func Select(data sort.Interface) {
	n := data.Len()
	if n <= 1 {
		return
	}
	for i := 0; i < n-1; i++ {
		minIndex := i
		for j := i + 1; j < n; j++ {
			if data.Less(minIndex, j) {
				minIndex = j
			}
		}
		if minIndex != i {
			data.Swap(i, minIndex)
		}
	}
}

// Insert 插入排序:从第一个元素开始，该元素可以认为已经被排序，取出下一个元素，
//  在已经排序的元素序列中从后向前扫描如果该元素（已排序）大于新元素，
//  将该元素移到下一位置，重复步骤3，直到找到已排序的元素小于或者等于新
//  元素的位置，将新元素插入到下一位置中，重复步骤2
// 平均时间复杂度：O(n^2)
// 最优时间复杂度：O(n)
// 最坏时间复杂度：O(n^2)
// 空间复杂度: O(1)
// 稳定性：YES
func Insert(data []int) {
	n := len(data)
	if n <= 1 {
		return
	}
	for i := 1; i < n; i++ {
		key, j := data[i], i-1
		for j >= 0 && data[j] > key {
			data[j+1] = data[j]
			j--
		}
		data[j+1] = key
	}
}

// Shell 排序
// 稳定性：NO
func Shell(data sort.Interface) {
	gap, n := 1, data.Len()
	if n <= 1 {
		return
	}

	for gap < n/3 {
		gap = gap*3 + 1
	}
	for gap > 0 {
		for i := gap; i < n; i++ {
			j := i
			for j >= gap && data.Less(j-gap, j) {
				data.Swap(j, j-gap)
				j = j - gap
			}
		}
		gap = gap / 3
	}
}

// Merge 归并排序
// 稳定性：YES
func Merge(data []int) []int {
	n := len(data)
	if n <= 1 {
		return data
	}
	// 分
	left := data[:n/2]
	lSize := len(left)
	if lSize > 1 {
		left = Merge(left)
	}

	right := data[n/2:]
	rSize := len(right)
	if rSize > 1 {
		right = Merge(right)
	}
	// 治
	i, j, result := 0, 0, make([]int, 0, n)
	for i < lSize || j < rSize {
		switch {
		case i < lSize && j < rSize:
			if left[i] <= right[j] {
				result = append(result, left[i])
				i++
			} else {
				result = append(result, right[j])
				j++
			}
		case i < lSize:
			result = append(result, left[i:]...)
			i += lSize - i
		case j < rSize:
			result = append(result, right[j:]...)
			j += rSize - j
		}
	}
	return result
}

func MergeByIterate(data []int) []int {
	n := len(data)
	if n <= 1 {
		return data
	}
	mid := n / 2
	return merge(MergeByIterate(data[:mid]), MergeByIterate(data[mid:]))
}

func merge(left, right []int) []int {
	lSize, rSize := len(left), len(right)
	i, j, result := 0, 0, make([]int, 0, len(left)+len(right))
	for i < lSize || j < rSize {
		switch {
		case i < lSize && j < rSize:
			if left[i] <= right[j] {
				result = append(result, left[i])
				i++
			} else {
				result = append(result, right[j])
				j++
			}
		case i < lSize:
			result = append(result, left[i:]...)
			i += lSize - i
		case j < rSize:
			result = append(result, right[j:]...)
			j += rSize - j
		}
	}
	return result
}

// Quick 快速排序
// 稳定性：NO
func Quick(data sort.Interface) {
	sort.Sort(data)
}

// Heap 堆排序
// 稳定性：NO
func Heap() {

}
