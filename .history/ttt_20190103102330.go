/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : ttt.go
#   Created       : 2019-01-03 10:22:53
#   Describe      :
#
# ====================================================*/
package main

import (
	"bytes"
	"strconv"
)

func main(){
	fmt.Println(maximumSwap(9683))
}

func maximumSwap(num int) int {
	if num < 10 {
		return num
	}
	max := num
	numStrs := strconv.Itoa(num)
	numStr := bytes.NewBufferString(numStrs).Bytes()
	start := 0
	for i := 0; i < len(numStr) && numStr[i] == '9'; i++ {
		start = i + 1
	}

	maxIdx := len(numStr) - 1
	maxN := numStr[maxIdx]
	for j := len(numStr) - 1; j >= start; j-- {
		if numStr[j] > maxN {
			maxIdx = j
			maxN = numStr[j]
		}
	}

	for j := start; j < len(numStr); j++ {
		if numStr[j] < maxN {
			numStr[j], numStr[maxIdx] = numStr[maxIdx], numStr[j]
			break
		}
	}

	tmp, _ := strconv.Atoi(string(numStr))
	if tmp > max {
		max = tmp
	}
	return max
}
