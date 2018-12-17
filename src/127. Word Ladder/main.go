package main

import (
	"container/list"
	"fmt"
)

//BFS
func ladderLength1(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 1
	}

	maps := make(map[string]bool)

	for _, word := range wordList {
		maps[word] = true
	}

	if _, ok := maps[endWord]; !ok {
		return 0
	}

	distance := 2

	len1 := len(beginWord)

	queue := list.New()
	queue.PushBack(beginWord)

	for queue.Len() != 0 {
		childQueue := list.New()
		for queue.Len() != 0 {
			curNode := queue.Front()
			queue.Remove(curNode)
			cur := curNode.Value.(string)
			delete(maps, cur)
			array := []byte(cur)
			for i := 0; i < len1; i++ {
				original := array[i]
				for j := 'a'; j <= 'z'; j++ {
					array[i] = byte(j)
					temp := string(array)
					if temp == endWord {
						return distance
					}

					if _, ok := maps[temp]; ok {
						childQueue.PushBack(temp)
						delete(maps, temp)
					}
				}
				array[i] = original
			}
		}
		queue = childQueue
		distance++
	}
	return 0
}

//double BFS
func ladderLength(beginWord string, endWord string, wordList []string) int {
	if beginWord == endWord {
		return 1
	}

	maps := make(map[string]bool)

	for _, word := range wordList {
		maps[word] = true
	}

	if _, ok := maps[endWord]; !ok {
		return 0
	}

	distance := 2

	src := map[string]bool{beginWord: true}
	dst := map[string]bool{endWord: true}

	for len(src) > 0 && len(dst) > 0 {
		if len(src) > len(dst) {
			src, dst = dst, src
		}

		toVisit := make(map[string]bool)

		for word, _ := range src {
			delete(maps, word)

			array := []byte(word)
			for i := 0; i < len(array); i++ {
				original := array[i]
				for j := 0; j < 26; j++ {
					array[i] = byte('a' + j)
					temp := string(array)
					if _, ok := dst[temp]; ok {
						return distance
					}
					if _, ok := maps[temp]; ok {
						toVisit [temp] = true
						delete(maps, temp)
					}
				}
				array[i] = original
			}
		}
		src = toVisit
		distance++
	}
	return 0
}

func canTransform(left string, right string) bool {
	len1 := len(left)
	count := 0
	for i := 0; i < len1; i++ {
		if left[i] != right[i] {
			count++
			if count > 1 {
				return false
			}
		}
	}
	return true
}

func main() {
	wordList := []string{"hot", "dot", "dog", "lot", "log", "cog"}
	fmt.Println(ladderLength1("hit", "cog", wordList))
}
