package main

import (
	"fmt"
	"sort"
)

func findLucky(arr []int) int {
	luckyNumber := -1

	sort.Ints(arr)

	largest := -1
	currentNumber := arr[0]
	count := 1
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1] != currentNumber {
			if count == currentNumber {
				if currentNumber > largest {
					largest = currentNumber
				}
			}
			count = 1
			currentNumber = arr[i+1]
		} else {
			count++
		}
	}
	if count == currentNumber {
		if currentNumber > largest {
			largest = currentNumber
		}
	}
	luckyNumber = largest

	return luckyNumber
}

func main() {
	fmt.Println(findLucky([]int{2, 2, 2, 3, 3}))
}
