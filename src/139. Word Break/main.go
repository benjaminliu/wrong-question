package main

import "fmt"

func wordBreak(s string, wordDict []string) bool {
	len1 := len(s)
	if len1 == 0 {
		return false
	}

	maps := make(map[string]bool)
	for _, value := range wordDict {
		maps[value] = true
	}

	len2 := len1 + 1
	dp := make([]bool, len2)
	dp[0] = true
	for i := 1; i < len2; i++ {
		for j := 0; j < i; j++ {
			if dp[j] {
				str := string(s[j:i])
				if _, ok := maps[str]; ok {
					dp[i] = true
					break
				}
			}
		}
	}
	return dp[len1]
}

//easier to understand
func wordBreak1(s string, wordDict []string) bool {
	slen := len(s)
	dp := make([]bool, slen+1)
	dp[0] = true
	for i := 1; i <= slen; i++ {
		for _, word := range wordDict {
			size := len(word)
			if i < size {
				continue
			}

			if dp[i-size] && s[i-size:i] == word {
				dp[i] = true
				break
			}
		}
	}

	return dp[slen]
}

//failed
func wordBreak2(s string, wordDict []string) bool {
	len1 := len(s)
	if len1 == 0 {
		return false
	}
	maps := make(map[string]bool)
	sizes := make(map[int]bool)
	for _, value := range wordDict {
		maps[value] = true
		sizes[len(value)] = true
	}

	return helper(s, 0, len1, maps, sizes)
}
func helper(s string, idx int, len1 int, maps map[string]bool, sizes map[int]bool) bool {
	if idx == len1 {
		return true
	}

	if idx > len1 {
		return false
	}

	for key, _ := range sizes {
		end := idx + key
		if end > len1 {
			continue
		}
		temp := string(s[idx:end])
		if _, ok := maps[temp]; ok {
			if helper(s, end, len1, maps, sizes) == true {
				return true
			}
		}
	}
	return false
}

func main() {
	s := "leetcode"
	wordList := []string{"leet", "code"}

	//s := "applepenapple"
	//wordList := []string{"apple", "pen"}

	//s := "catsandog"
	//wordList := []string{"cats", "dog", "sand", "and", "cat"}

	//s := "bb"
	//wordList := []string{"a", "b", "bbb", "bbbb"}

	//s := "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"
	//wordList := []string{"a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"}

	fmt.Println(wordBreak(s, wordList))
}
