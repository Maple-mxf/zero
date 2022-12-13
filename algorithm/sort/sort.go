
package sort

import (
	"fmt"
)

type Num interface {
	int32 | int64 | float32 | float64
}

// 快速排序算法定义：选择一个基准元素pivot，定义两个指针，left和right
// left指向第一个元素，right指向第二个元素，从right指针开始，
// 如果right指针指向的元素小于基准元素，·1
// QuickSort 快速排序算法
func QuickSort[T Num](nums []T) {
	_quickSort(nums, 0, len(nums)-1)
}

func _quickSort[T Num](nums []T, startIndex, endIndex int) {
	left, right := startIndex, endIndex
	if left < right {
		pivot := nums[left]
		for left != right {

			for right > left && nums[right] >= pivot {
				right--
			}
			nums[left] = nums[right]

			for left < right && nums[left] <= pivot {
				left++
			}
			nums[right] = nums[left]
		}

		nums[right] = pivot
		_quickSort(nums, startIndex, left-1)
		_quickSort(nums, right+1, endIndex)
	}
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
