/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 150.evaluate_reverse_polish_notation.go
#   Created       : 2019-02-12 16:42:37
#   Last Modified : 2019-02-12 16:42:37
#   Describe      :
#
# ====================================================*/
package leet

import "strconv"

func evalRPN(tokens []string) int {
	var stack []int
	for _, v := range tokens {
		if v != "+" && v != "/" && v != "*" && v != "-" {
			cv, _ := strconv.Atoi(v)
			stack = append(stack, cv)
		} else {
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			r := getRes(v, a, b)
			stack = append(stack, r)
		}
	}
	return stack[0]
}

func getRes(v string, a, b int) int {
	switch v {
	case "+":
		return a + b
	case "/":
		return a / b
	case "-":
		return a - b
	case "*":
		return a * b
	}
	return 0
}
