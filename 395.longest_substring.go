/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 395.longest_substring.go
#   Created       : 2019-03-13 14:08:04
#   Last Modified : 2019-03-13 14:08:04
#   Describe      :
#
# ====================================================*/
package main

func longestSubstring(s string, k int) int {
	if len(s) == 0 || k == 1 {
		return len(s)
	}

	more := make(map[byte]int)
	less := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		if more[s[i]] >= k {
			more[s[i]]++
		} else {
			less[s[i]]++
			if less[s[i]] >= k {
				more[s[i]] = less[s[i]]
				delete(less, s[i])
			}
		}
	}
	if len(less) == 0 {
		return len(s)
	}

	start := 0
	ret := 0
	for i := 0; i < len(s); i++ {
		if _, ok := less[s[i]]; ok {
			ret = max(ret, longestSubstring(s[start:i], k))
			start = i + 1
		}
	}
	return max(ret, longestSubstring(s[start:], k))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
