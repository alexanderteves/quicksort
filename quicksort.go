package quicksort

import (
	"math/rand"
	"time"
)

func Sort(sortList []int) []int {
	if len(sortList) > 1 {
		if len(sortList) == 2 {
			a, b := sortList[0], sortList[1]
			if a == b || a < b {
				return sortList
			} else {
				return []int{b, a}
			}
		}
		rand.Seed(time.Now().Unix())
		pivot := rand.Intn(len(sortList))

		lower := make([]int, 0)
		equal := make([]int, 0)
		greater := make([]int, 0)

		for _, v := range sortList {
			if v < sortList[pivot] {
				lower = append(lower, v)
			} else if v > sortList[pivot] {
				greater = append(greater, v)
			} else {
				equal = append(equal, v)
			}
		}

		lower = Sort(lower)
		greater = Sort(greater)

		sortList = append(lower, equal...)
		sortList = append(sortList, greater...)
	}
	return sortList
}
