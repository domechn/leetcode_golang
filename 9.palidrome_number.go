// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package main

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	com := x
	newX := 0
	for x != 0 {
		newX = newX*10 + x%10
		x = x / 10
	}
	return newX == com
}
