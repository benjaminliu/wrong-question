package main

import (
	"fmt"
	"strconv"
)

/*
   没做出来！！！
   思路：本题递归实现
   采用分治算法，分治算法的基本思想是将一个规模为N的问题分解为K个规模较小的子问题，
   这些子问题相互独立且与原问题性质相同，求出子问题的解，就可得到原问题的解。
   那么针对本题，以操作符为分界，将字符串分解为较小的两个子字符串，然后依次对两个子字符串进行同样的划分，
   直到字符串中只含有数字。再根据操作符对两端的数字进行相应的运算。
   */

var maps map[string][]int = make(map[string][]int)

func diffWaysToCompute(input string) []int {
if value, ok := maps[input]; ok {
	return value
}

len1 := len(input)

res := make([]int, 0)
for i := 0; i < len1; i++ {
	if input[i] == '+' || input[i] == '-' || input[i] == '*' {
		left := diffWaysToCompute(string(input[0:i]))
		right := diffWaysToCompute(string(input[i+1:]))

		for _, leftValue := range left {
			for _, rightValue := range right {
				if input[i] == '+' {
					res = append(res, leftValue+rightValue)
				}
				if input[i] == '-' {
					res = append(res, leftValue-rightValue)
				}
				if input[i] == '*' {
					res = append(res, leftValue*rightValue)
				}
			}
		}
	}
}

if len(res) == 0 {
	value, _ := strconv.Atoi(input)
	res = append(res, value)
}

maps[input] = res

return res
}

func main() {
input := "2*3-4*5"
fmt.Println(diffWaysToCompute(input))
}
