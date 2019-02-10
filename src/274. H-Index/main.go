package main

func hIndex(citations []int) int {
	len1 := len(citations)
	buckets := make([]int, len1+1)
	for _, value := range citations {
		if value > len1 {
			buckets[len1]++
		} else {
			buckets[value]++
		}
	}

	count := 0
	for i := len1; i >= 0; i-- {
		count += buckets[i]
		if count >= i {
			return i
		}
	}
	return 0
}

func main() {

}
