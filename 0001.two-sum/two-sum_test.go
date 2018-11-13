package twosum

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	result := TwoSum([]int{2, 7, 11, 15}, 9)
	if len(result) != 2 || result[0] != 0 || result[1] != 1 {
		t.Error("结果错误:", result)
	}
}
