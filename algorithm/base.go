package algorithm

var TowNumsSumCount = 0

// TowNumsSum 求两数字之和，入参数组里面的数字都是不重复的
func TowNumsSum(nums []int, result int) (ret []int) {
	for i := 0; i < len(nums); i++ {
		var diff = result - nums[i]
		for k := i + 1; k < len(nums); k++ {
			if nums[k] == diff {
				ret = append(ret, nums[i])
				ret = append(ret, nums[k])
				return
			}
		}
	}
	return
}

func TestNumsSum2(nums []int, result int) (ret []int) {
	var m = make(map[int]interface{})
	for _, num := range nums {
		if _, ok := m[result-num]; ok {
			ret = append(ret, result-num)
			ret = append(ret, num)
			return
		}
		m[num] = 1
	}
	return
}
