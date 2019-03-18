/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 1002.find_common_characters.go
#   Created       : 2019-03-04 15:34:28
#   Last Modified : 2019-03-04 15:34:28
#   Describe      :
#
# ====================================================*/
package main

func commonChars(A []string) []string {
	if len(A) == 0 {
		return []string{}
	}

	var temp = make(map[int32]int)

	for _, v := range A[0] {
		temp[v]++
	}

	for i := 1; i < len(A); i++ {
		var tem2 = make(map[int32]int)
		for _, v := range A[i] {
			tem2[v]++
		}
		for k, v := range temp {
			if kk, ok := tem2[k]; ok {
				temp[k] = min(v, kk)
			} else {
				delete(temp, k)
			}
		}
	}
	var res []string
	for k, v := range temp {
		for i := 0; i < v; i++ {
			res = append(res, string(k))
		}
	}
	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
