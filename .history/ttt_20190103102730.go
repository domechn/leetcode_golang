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
	tmap := make(map[int]int,10)
	i := 1
	for num > 0 {
		tmap[i] = num%10
		i *= 10
		num /= 10
	}
	for j := 1; j <= i && tmap[j]==9; j*=10 {
		start = j * 10
	}

	maxIdx := i
	maxN := tmap[i]
	for j := ; j >= start; j /=10 {
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
