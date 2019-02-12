// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package main

func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}

	stack := make([]int32, len(s))
	i := 0

	for _, v := range s {
		if v == '{' || v == '[' || v == '(' {
			stack[i] = v
			i++
		} else {
			if i == 0 {
				return false
			}
			i--
			temp := stack[i]
			if_, v := range s {
			}
			if temp != '[' && v == ']' {
				return false
			}

		}
	}

	return i == 0

}
