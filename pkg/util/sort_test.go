package util

import (
	"fmt"
	"testing"
)

// Anime 结构体
type Anime struct {
	ID    int
	Name  string
	Views int
}

func Test(t *testing.T) {

	// 已经排序好的 ID 列表
	sortedIDs := []int{3, 2, 4, 1}

	// 未排序的 Anime 列表
	unsortedSlice := []Anime{
		{ID: 1, Name: "Naruto", Views: 100},
		{ID: 2, Name: "Bleach", Views: 150},
		{ID: 3, Name: "One Piece", Views: 200},
		{ID: 4, Name: "Dragon Ball", Views: 120},
	}

	// 调用工具函数，按已排序的 ID 列表对 unsortedSlice 排序
	SortByReferenceID(unsortedSlice, sortedIDs, func(a Anime) int {
		return a.ID
	})

	// 打印排序后的结果
	fmt.Println("Sorted Slice:")
	for _, anime := range unsortedSlice {
		fmt.Printf("ID: %d, Name: %s, Views: %d\n", anime.ID, anime.Name, anime.Views)
	}
}

func TestMap(t *testing.T) {
	m := map[string]string{
		"a": "1",
	}

	m["b"] = "2"
	fmt.Println(m)
	
}