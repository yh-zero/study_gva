package main

import (
	"fmt"
	"math/rand"
	"time"
)

func weightedLottery() int {
	// 设置随机数种子
	rand.Seed(time.Now().UnixNano())

	// 定义数字范围
	minNumber := 1
	maxNumber := 100

	// 定义中奖概率
	winProbabilities := make(map[int]int)
	for i := maxNumber; i >= minNumber; i-- {
		// 根据规则设定中奖概率，可以根据需要进行调整
		winProbabilities[i] = maxNumber - i + 1
	}

	// 计算总权重
	totalWeight := 0
	for _, probability := range winProbabilities {
		totalWeight += probability
	}

	// 生成一个随机数，用来确定抽中的数字
	randomValue := rand.Intn(totalWeight)

	// 遍历数字范围，根据权重确定抽中的数字
	cumulativeWeight := 0
	var selectedNumber int

	for i := minNumber; i <= maxNumber; i++ {
		if probability, ok := winProbabilities[i]; ok {
			cumulativeWeight += probability
			if randomValue < cumulativeWeight {
				selectedNumber = i
				break
			}
		}
	}

	return selectedNumber
}

func main() {
	// 调用抽奖函数
	winner := weightedLottery()
	fmt.Printf("恭喜，你中奖了！抽中的数字是：%d\n", winner)
}
