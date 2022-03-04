package sortpkg

//从小到大排序
// BubbleSort 冒泡排序，每次找出最小的值放在最前面
func BubbleSort(data []int) {
	count := len(data)
	if count <= 1 {
		return
	}

	for i := 0; i < count; i++ {
		for j := 0; j < count-i-1; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
}

// QuickSort 快速排序
func QuickSort(data []int) []int {
	count := len(data)
	if count <= 0 {
		return nil
	}

	//选择中间的数作为参考
	keyIndex := count / 2
	key := data[keyIndex]

	//分成左右两部分，左边放比key小的值，右边放比key大的值
	left := make([]int, 0)
	right := make([]int, 0)

	for i := 0; i < count; i++ {
		if i == keyIndex {
			continue
		}
		if data[i] < key {
			left = append(left, data[i])
		} else {
			right = append(right, data[i])
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)

	//最后将得到的两组数组合起来
	var result []int
	result = append(result, left...)
	result = append(result, key)
	result = append(result, right...)
	return result
}

// SelectSort 选择排序:每次选择出未排序切片里最大或者最小的数放入已排好序的数组里
func SelectSort(data []int) []int {
	count := len(data)
	if count <= 0 {
		return nil
	}

	var min, minIndex int
	for i := 0; i < count-1; i++ {
		min = data[i]
		minIndex = i
		for j := i + 1; j < count; j++ {
			if data[j] < min {
				min = data[j]
				minIndex = j
			}
		}
		data[i], data[minIndex] = data[minIndex], data[i]
	}
	return data
}

// InsertSort 插入排序:从第一个元素开始，该元素可以认为已经被排序，取出下一个元素，
//  在已经排序的元素序列中从后向前扫描如果该元素（已排序）大于新元素，
//  将该元素移到下一位置，重复步骤3，直到找到已排序的元素小于或者等于新
//  元素的位置，将新元素插入到下一位置中，重复步骤2
func InsertSort(data []int) {
	count := len(data)
	if count <= 1 {
		return
	}

	var key, pos int
	for i := 1; i < count; i++ {
		key = data[i]
		pos = i

		//此处循环为了将比key大的数往后移动
		for pos > 0 && data[pos-1] > key {
			data[pos] = data[pos-1]
			pos--
		}
		data[pos] = key
	}
}
