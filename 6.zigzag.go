// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package main

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}
	size := numRows*2 - 2
	res := ""
	for i := 0; i < numRows; i++ {
		for j := i; j < len(s); j += size {
			res += string(s[j])
			tmp := j + size - 2*i
			if i != 0 && i != numRows-1 && tmp < len(s) {
				res += string(s[tmp])
			}
		}
	}
	return res
}
