// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package main

func romanToInt(s string) int {
	var temp = map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	var res int
	for i := len(s) - 1; i >= 0; i-- {
		if i < len(s)-1 && temp[string(s[i+1])] > temp[string(s[i])] {
			res -= temp[string(s[i])]
		} else {
			res += temp[string(s[i])]
		}
	}
	return res
}
