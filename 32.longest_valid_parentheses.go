/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 32.longest_valid_parentheses.go
#   Created       : 2019-02-12 17:09:44
#   Last Modified : 2019-02-12 17:09:44
#   Describe      :
#
# ====================================================*/
package leet

func longestValidParentheses(s string) int {
	var stack []int32
	var res int
	var temp int
	for _, v := range s {
		if v == '(' {
			push(stack, v)
			continue
		}
		if len(stack) == 0 {
			temp = 0
			continue
		}
		t := pop(stack)
		if t == '(' {
			temp++
		}
		if res < temp {
			res = temp
		}
	}
	return res
}

func pop(stack []int32) int32 {
	s := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return s
}

func push(stack []int32, s int32) {
	stack = append(stack, s)
}
