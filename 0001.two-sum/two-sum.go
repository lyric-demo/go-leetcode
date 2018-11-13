/*
Package twosum 考察知识点：
1. 学会灵活使用字典
2. 减少遍历次数
*/
package twosum

// TwoSum 给定一个整数数组，返回两个数字的索引，使它们相加到特定目标。
func TwoSum(vals []int, target int) []int {
	m := make(map[int]int)

	for i := 0; i < len(vals); i++ {
		if v, ok := m[target-vals[i]]; ok {
			return []int{v, i}
		}
		m[vals[i]] = i
	}
	return nil
}
