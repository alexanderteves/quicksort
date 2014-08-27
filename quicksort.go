package quicksort

import (
	"math/rand"
	"time"
)

func Sort(sortList []int) []int {
	if len(sortList) > 1 {
		if len(sortList) == 2 {
			if a, b := sortList[0], sortList[1]; a == b || a < b {
				return sortList
			} else {
				return []int{b, a}
			}
		}
		rand.Seed(time.Now().Unix())
		pivot := rand.Intn(len(sortList))

		lowerValues := make([]int, 0, len(sortList))
		equalValues := make([]int, 0, len(sortList))
		greaterValues := make([]int, 0, len(sortList))

		for _, value := range sortList {
			if value < sortList[pivot] {
				lowerValues = append(lowerValues, value)
			} else if value > sortList[pivot] {
				greaterValues = append(greaterValues, value)
			} else {
				equalValues = append(equalValues, value)
			}
		}

		lowerValues = Sort(lowerValues)
		greaterValues = Sort(greaterValues)

		sortList = append(lowerValues, equalValues...)
		sortList = append(sortList, greaterValues...)
	}
	return sortList
}
