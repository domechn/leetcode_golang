// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package main

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}

	if needle == haystack {
		return 0
	}

	for i := 0; i < len(haystack)-len(needle)+1; i++ {
		if haystack[i] == needle[0] && haystack[i:i+len(needle)] == needle {
			return i
		}
	}
	return -1
}
