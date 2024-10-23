package util

import "sort"

/*
SortByReferenceID 对未排序的 slice 按参考 ID 排序

参数1：未排序的 slice

参数2：已排序的 ID 列表

参数3：一个函数，用于从 slice 中提取 ID
*/
func SortByReferenceID[T any](slice []T, sortedIDs []int, getID func(item T) int) {
	// 构建 ID 到索引的映射表
	idIndexMap := make(map[int]int)
	for index, id := range sortedIDs {
		idIndexMap[id] = index
	}

	// 使用 sort.Slice 来对未排序的 slice 进行排序
	sort.Slice(slice, func(i, j int) bool {
		return idIndexMap[getID(slice[i])] < idIndexMap[getID(slice[j])]
	})
}
