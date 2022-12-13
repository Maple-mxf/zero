package sort

import (
	"fmt"
	"math"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	var arr = []float64{10, 1000, 2, 100, 1}
	fmt.Println(arr)
	BubbleSort(arr)
	fmt.Println(arr)
}

func TestQuickSort(t *testing.T) {
	var arr = []float64{10, 1000, 2, 100, 1, 3}
	fmt.Println(arr)
	QuickSort(arr)
	fmt.Println(arr)
}

func TestMathLog(t *testing.T) {
	fmt.Println(math.Log10(5))
	fmt.Println(math.Log(5))
}
