package sorts

import (
	"sort"
)

// BubbleSort 冒泡排序，每次找出最小的值放在最前面
func BubbleSort(data sort.Interface) {
	existSwap := false // 是否有发生了交换动作, 如果没有发生交换，证明数据已经有序了, 不需要在进行遍历比较
	for i := 0; i < data.Len()-1; i++ {
		temp := data.Len() - 1 - i
		for j := 0; j < temp; j++ {
			if data.Less(j+1, j) {
				data.Swap(j, j+1)
				existSwap = true
			}
		}
		if !existSwap {
			break
		}
	}
}

// QuickSort 快速排序
func QuickSort(data sort.Interface) {
	sort.Sort(data)
}

// SelectSort 选择排序:每次选择出未排序切片里最大或者最小的数放入已排好序的数组里
func SelectSort(data sort.Interface) {
	count := data.Len()
	for i := 0; i < count-1; i++ {
		minIndex := i
		for j := i + 1; j < count; j++ {
			if data.Less(j, minIndex) {
				minIndex = j
			}
		}
		data.Swap(i, minIndex)
	}
}

// InsertSort 插入排序:从第一个元素开始，该元素可以认为已经被排序，取出下一个元素，
//  在已经排序的元素序列中从后向前扫描如果该元素（已排序）大于新元素，
//  将该元素移到下一位置，重复步骤3，直到找到已排序的元素小于或者等于新
//  元素的位置，将新元素插入到下一位置中，重复步骤2
func InsertSort(data sort.Interface) {
	count := data.Len()
	for i := 1; i < count; i++ {
		for j := i; j > 0; j-- {
			if data.Less(j, j-1) {
				data.Swap(j, j-1)
			}
		}
	}
}
