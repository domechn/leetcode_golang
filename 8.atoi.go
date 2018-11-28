// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package main

import (
	"fmt"
)

const MaxInt32 = 1<<31 - 1
const MinInt32 = -1 << 31

func main() {
	fmt.Println(myAtoi("0-1"))
}

func myAtoi(str string) int {
	var res int
	var flag = 1
	var judgePositive, judgeBalnk bool
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' && !judgeBalnk {
			continue
		}
		judgeBalnk = true
		if str[i] == '-' && !judgePositive {
			flag = -1
			judgePositive = true
			continue
		} else if str[i] == '+' && !judgePositive {
			flag = 1
			judgePositive = true
			continue
		}
		judgePositive = true
		if str[i] < '0' || str[i] > '9' {
			break
		}
		res = res*10 + int(str[i]-'0')
		if res*flag > MaxInt32 {
			return MaxInt32
		}
		if res*flag < MinInt32 {
			return MinInt32
		}
	}

	return res * flag
}
