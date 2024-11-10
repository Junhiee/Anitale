package util

import (
	"math"
	"time"
)

type Stats struct {
	ViewCount    int64     // 播放数量
	LikeCount    int64     // 点赞数量
	CommentCount int64     // 评论数量
	ShareCount   int64     // 转发数量
	UpdatedAt    time.Time // 更新时间
	CreatedAt    time.Time // 创建时间
}

// 计算热度值的函数，加入时间衰减因素
func CalculateHeatValue(stats Stats, currentTime time.Time) float64 {
	// 权重设置（可根据实际情况调整）
	viewWeight := 0.1
	likeWeight := 0.2
	commentWeight := 0.4
	shareWeight := 0.7

	// 设定半衰期（7天）
	halfLife := 7 * 24 * 60 * 60 // 7天的秒数

	// 计算时间差（当前时间 - 创建时间），单位为秒
	timeDifference := currentTime.Sub(stats.UpdatedAt).Seconds()

	// 限制最大时间差为一年的秒数，避免衰减因子过小
	maxTimeDifference := 365 * 24 * 60 * 60 // 1年的秒数
	if timeDifference > float64(maxTimeDifference) {
		timeDifference = float64(maxTimeDifference)
	}

	// 计算时间衰减因子
	decayFactor := math.Exp(-timeDifference / float64(halfLife))

	// 根据权重和衰减因子计算热度值
	heatValue := decayFactor * (float64(stats.ViewCount)*viewWeight +
		float64(stats.LikeCount)*likeWeight +
		float64(stats.CommentCount)*commentWeight +
		float64(stats.ShareCount)*shareWeight)

	return heatValue
}
