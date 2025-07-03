package main

import "fmt"

func main() {
	fmt.Println(kthCharacter(10))
	fmt.Println(nextChar('z'))
}

func kthCharacter(k int) byte {
	word := "a"
	wordBytes := []byte(word)
	newWordBytes := []byte(word)
	for {
		for i := 0; i < len(wordBytes); i++ {
			newWordBytes = append(newWordBytes, nextChar(wordBytes[i]))
		}
		wordBytes = newWordBytes
		if len(newWordBytes) >= k {
			break
		}
	}

	return newWordBytes[k-1]
}

func nextChar(ch byte) byte {
	if ch += 1; ch > 'z' {
		return 'a'
	}
	return ch
}
