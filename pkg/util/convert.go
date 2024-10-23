package util

import "strconv"

// ConvertStringsToInts 将 []string 转换为 []int
func ConvertStringsToInts(stringSlice []string) ([]int, error) {
	intSlice := make([]int, len(stringSlice)) // 创建一个与输入 stringSlice 长度相同的 int 切片
	for i, s := range stringSlice {
		num, err := strconv.Atoi(s) // 将字符串转换为整数
		if err != nil {
			return nil, err // 如果转换失败，返回错误
		}
		intSlice[i] = num
	}
	return intSlice, nil
}
