package main

import (
	"fmt"
	"sort"
)

func maxEvents(events [][]int) int {
	maxDays := 0
	sort.Slice(events, func(i, j int) bool {
		if events[i][1] > maxDays {
			maxDays = events[i][1]
		}
		return events[i][0] < events[j][0]
	})

	sort.SliceStable(events, func(i, j int) bool {
		return events[i][1] < events[j][1]
	})

	if maxDays == len(events) {
		return maxDays
	}

	reservedDays := make(map[int]bool, maxDays)
	lastDay := 1
	for i := 0; i < len(events); i++ {
		for j := lastDay; j <= events[i][1]; j++ {
			if !reservedDays[j] {
				reservedDays[j] = true
				lastDay = j
				break
			}
		}
	}
	return len(reservedDays)
}

func main() {
	fmt.Println(maxEvents([][]int{{1, 2}, {1, 2}, {1, 6}, {1, 2}, {1, 2}}))
}
