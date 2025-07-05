/*package main

import "fmt"

func main() {
	fmt.Println(kthCharacter(5, []int{0, 0, 0}))
}

func kthCharacter(k int, operations []int) byte {
	word := "a"
	wordBytes := []byte(word)
	newWordBytes := []byte(word)

	for i := 0; i < len(operations); i++ {
		if operations[i] == 0 {
			//newWordBytes = append(newWordBytes, newWordBytes...)
			k /= 2
		}
		if operations[i] == 1 {
			for i := 0; i < len(wordBytes); i++ {
				newWordBytes = append(newWordBytes, nextChar(wordBytes[i]))
			}
		}
		wordBytes = newWordBytes
	}

	return newWordBytes[k-1]
}

func nextChar(ch byte) byte {
	if ch += 1; ch > 'z' {
		return 'a'
	}
	return ch
}
*/

package main

import (
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(string(kthCharacter(33354182522397, []int{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0, 1, 1, 0, 1, 1, 1, 0, 1, 0, 1, 0, 1, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 1, 0, 0, 0, 1, 0, 1, 1, 0, 1, 0, 0, 0, 1, 0, 1, 0, 1, 1, 0, 0, 0, 0, 1, 0, 1, 1, 0, 0, 1, 0, 0, 1, 1, 0, 1, 1, 1, 1, 1, 1, 0, 0, 0})))
}

func kthCharacter(k int64, operations []int) byte {
	ans := 0
	for k != 1 {
		t := bits.Len64(uint64(k)) - 1
		if (1 << t) == k {
			t--
		}
		k -= 1 << t
		if operations[t] != 0 {
			ans++
		}
	}
	return byte('a' + (ans % 26))
}
