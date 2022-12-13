package algorithm

import (
	"fmt"
	"testing"
)

func TestTowNumsSum(t *testing.T) {
	ret := TowNumsSum([]int{1, 3, 2, 7}, 10)
	fmt.Println(ret)
	fmt.Println(TowNumsSumCount)
}

func TestTowNumsSum2(t *testing.T) {
	ret := TestNumsSum2([]int{1, 3, 2, 7}, 10)
	fmt.Println(ret)
	fmt.Println(TowNumsSumCount)
}
