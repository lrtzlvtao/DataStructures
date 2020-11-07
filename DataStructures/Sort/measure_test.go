package sort

import (
	"testing"
)

// 随机数排序比较
func TestMeasure(t *testing.T) {
	l := randomList(10000)
	insertionTime := measure("Insertion", l)
	selectionTime := measure("Selection", l)
	shellTime := measure("Shell", l)
	t.Log(insertionTime)
	t.Log(selectionTime)
	t.Log(shellTime)
}

// 有序数据比较
func TestMeasure2(t *testing.T) {
	l := randomSortedList(10000)
	insertionTime := measure("Insertion", l)
	selectionTime := measure("Selection", l)
	shellTime := measure("Shell", l)
	t.Log(insertionTime)
	t.Log(selectionTime)
	t.Log(shellTime)
}
