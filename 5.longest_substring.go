// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package main

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}
	var res string
	for i := 0; i < len(s)-1; i++ {
		for j := len(s) - 1; j >= i; j-- {
			if s[i] == s[j] {
				jg := judge(s[i : j+1])
				if len(jg) > len(res) {
					res = jg
				}
			}
		}
	}
	return res
}

func judge(s string) string {
	if len(s) == 1 {
		return s
	}
	begin := 0
	end := len(s)
	for {
		if begin >= end {
			break
		}
		if s[begin] == s[end-1] {
			begin++
			end--
		} else {
			return ""
		}
	}
	return s
}
