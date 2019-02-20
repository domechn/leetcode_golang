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

import "fmt"

func main() {
	fmt.Println(maximumSwap(9683))
}

func maximumSwap(num int) int {
	max := num
	if num < 10 {
		return num
	}
	tmap := make(map[int]int, 10)
	start := 0
	i := 1
	for num > 0 {
		tmap[i] = num % 10
		i *= 10
		num /= 10
	}
	fmt.Println(tmap)
	for j := i; j <= i && tmap[j] == 9; j /= 10 {
		start = j * 10
	}

	maxIdx := i
	maxN := tmap[i]
	for j := start; j >= 1; j /= 10 {
		if tmap[j] > maxN {
			maxIdx = j
			maxN = tmap[j]
		}
		if j == 1 {
			break
		}
	}
	fmt.Println(start)

	for j := start; j >= 1; j /= 10 {
		if tmap[j] < maxN {
			tmap[j], tmap[maxIdx] = tmap[maxIdx], tmap[j]
			break
		}
		if j == 1 {
			break
		}
	}

	tmp := 0
	for k, v := range tmap {
		tmp += k * v
	}

	if tmp > max {
		max = tmp
	}
	return max
}
