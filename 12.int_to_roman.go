// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package main

func intToRoman(num int) string {
	numList := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	strList := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	var res string
	for k, v := range numList {
		for num-v >= 0 {
			res += strList[k]
			num -= v
		}
	}
	return res
}
