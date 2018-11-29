// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	res := strs[0]
	for i := 1; i < len(strs); i++ {
		res = prefix(res, strs[i])
	}
	return res
}

func prefix(s1, s2 string) string {
	if s1 == "" || s2 == "" {
		return ""
	}
	i := 0

	for i < len(s1) && i < len(s2) {
		if s1[i] == s2[i] {
			i++
		} else {
			break
		}
	}
	return s1[:i]
}
