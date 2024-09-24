package main

import "fmt"

func main() {
	fmt.Println(Ft_max_substring("abcabcbb")) // resultat : 3
	// "abc" est la plus grande sous chaine composé de caractère diffèrent
	fmt.Println(Ft_max_substring("bbbbb")) // resultat : 1
	// "b" est la plus grande sous chaine
}

func Ft_max_substring(s string) int {
	if len(s) == 0 {
		return 0
	}

	maxLen := 1
	start := 0
	charMap := make(map[rune]int)

	for i, char := range s {
		if prevIndex, ok := charMap[char]; ok && prevIndex >= start {
			start = prevIndex + 1
		} else {
			maxLen = max(maxLen, i-start+1)
		}
		charMap[char] = i
	}

	return maxLen
}
