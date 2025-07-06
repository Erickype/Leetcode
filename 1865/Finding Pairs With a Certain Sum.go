package main

import "fmt"

type FindSumPairs struct {
	Nums1     []int
	Nums2     []int
	FreqNums2 map[int]int // frequency map for Nums2
}

func Constructor(nums1 []int, nums2 []int) FindSumPairs {
	freq := make(map[int]int)
	for _, num := range nums2 {
		freq[num]++
	}
	return FindSumPairs{
		Nums1:     nums1,
		Nums2:     nums2,
		FreqNums2: freq,
	}
}

func (this *FindSumPairs) Add(index int, val int) {
	oldVal := this.Nums2[index]
	this.FreqNums2[oldVal]--
	if this.FreqNums2[oldVal] == 0 {
		delete(this.FreqNums2, oldVal)
	}

	this.Nums2[index] += val
	newVal := this.Nums2[index]
	this.FreqNums2[newVal]++
}

func (this *FindSumPairs) Count(tot int) int {
	count := 0
	for _, a := range this.Nums1 {
		b := tot - a
		if freq, ok := this.FreqNums2[b]; ok {
			count += freq
		}
	}
	return count
}

func main() {
	findSumPairs := Constructor([]int{1, 1, 2, 2, 2, 3}, []int{1, 4, 5, 2, 5, 4})
	findSumPairs.Add(3, 2)             // nums2 = [1,4,5,4,5,4]
	fmt.Println(findSumPairs.Count(8)) // Expected: 2
	fmt.Println(findSumPairs.Count(4)) // Expected: 1
	findSumPairs.Add(0, 1)             // nums2 = [2,4,5,4,5,4]
	findSumPairs.Add(1, 1)             // nums2 = [2,5,5,4,5,4]
	fmt.Println(findSumPairs.Count(7)) // Output depends on nums1 & nums2
}
