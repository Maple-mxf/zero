package sort

import (
	"fmt"
)

type Num interface {
	int32 | int64 | float32 | float64
}

// 快速排序算法定义：选择一个基准元素pivot，定义两个指针，left和right
// left指向第一个元素，right指向第二个元素，从right指针开始，
// 如果right指针指向的元素小于基准元素，则替换
// QuickSort 快速排序算法
func QuickSort[T Num](nums []T) {
	_quickSort(nums, 0, len(nums)-1)
}

func _quickSort[T Num](nums []T, startIndex, endIndex int) {
	if startIndex >= endIndex {
		return
	}
	pivotIndex := partition(nums, startIndex, endIndex)
	_quickSort(nums, startIndex, pivotIndex-1)
	_quickSort(nums, pivotIndex+1, endIndex)
}

func partition[T Num](nums []T, startIndex, endIndex int) (index int) {
	pivot := nums[0]
	left, right := startIndex, endIndex
	index = startIndex

	for right > left {

		// 从右侧向左侧遍历，直到遇到右侧的
		for right > left {
			if nums[right] < pivot {
				nums[left] = nums[right]
				index = right
				left++
				break
			}
			right--
		}

		for right > left {
			if nums[left] > pivot {
				nums[right] = nums[left]
				index = left
				right--
				break
			}
			left++
		}

	}
	nums[index] = pivot
	return
}

// 它重复地走访过要排序的元素列，依次比较两个相邻的元素，
// 如果顺序（如从大到小、首字母从Z到A）错误就把他们交换过来。
// 走访元素的工作是重复地进行，直到没有相邻元素需要交换，
// 也就是说该元素列已经排序完成。
// 算法复杂度分析

// BubbleSort 冒泡排序算法
func BubbleSort[T Num](nums []T) {
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			count++
			if nums[j] > nums[j+1] {
				temp := nums[j]
				nums[j] = nums[j+1]
				nums[j+1] = temp
			}
		}
	}
	fmt.Println(count)
}
