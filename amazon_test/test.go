package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

/*
 * Complete the 'orderBooks' function below.
 *
 * The function is expected to return a 2D_INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY volumes as parameter.
 */

func orderBooks(volumes []int32) [][]int32 {
	// Write your code here
	volumesPurchased := make([][]int32, len(volumes))
	availableVolumes := make([]int32, len(volumes))
	for i := 0; i < len(volumes); i++ {
		availableVolumes[i] = volumes[i]
		if isVolumePurchased(volumesPurchased[i], volumes[i]) {
			noPurchase := []int32{-1}
			volumesPurchased[i] = noPurchase
		} else {
			if arePrecuelsAvailable(availableVolumes, volumes[i]) {
				volumesPurchased[i] = buyNotOwned(volumesPurchased, availableVolumes)
			} else {
				noPurchase := []int32{-1}
				volumesPurchased[i] = noPurchase
			}
		}
	}
	return volumesPurchased
}

func buyNotOwned(purchased [][]int32, volumes []int32) []int32 {
	newVolumes := []int32{}
	for i := 0; i < len(volumes); i++ {
		for j := 0; j < len(purchased); j++ {
			for k := 0; k < len(purchased[j]); k++ {
				if purchased[j][k] != volumes[i] {
					newVolumes = append(newVolumes, purchased[j][k])
				}
			}
		}
	}
	return newVolumes
}

func isVolumePurchased(purchases []int32, newVolume int32) bool {
	for i := 0; i < len(purchases); i++ {
		if purchases[i] == newVolume {
			return true
		}
	}
	return false
}

func arePrecuelsAvailable(availableVolumes []int32, volumeToPurchase int32) bool {
	sort.Slice(availableVolumes, func(i, j int) bool {
		return availableVolumes[i] < availableVolumes[j]
	})

	if len(availableVolumes) < int(volumeToPurchase) {
		return false
	}
	for i := 0; i <= int(volumeToPurchase); i++ {
		for _, volume := range availableVolumes {
			if volume != int32(i) && volume != 0 {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println(orderBooks([]int32{2, 1, 4, 3}))
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
