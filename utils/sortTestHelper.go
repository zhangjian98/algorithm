package utils

import (
	"math/rand"
	"time"
)

//生成n个元素的随机数组,每个随机数组范围[rangL, rangR]
func GenerateRandomArray(n int, rangeL int, rangeR int) []int {
	//范围检查
	if rangeR < rangeL || (rangeR - rangeL) > n {
		return nil
	}

	arr := make([]int,0)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(arr) < n {
		arr = append(arr, r.Intn((rangeR - rangeL + 1) + rangeL))
	}

	return arr
}