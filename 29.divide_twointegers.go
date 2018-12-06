// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package main

const MaxInt32 = 1<<31 - 1
const MinInt32 = -1 << 31

func divide(dividend int, divisor int) int {
	if divisor == 0 || dividend == MinInt32 && divisor == -1 {
		return MaxInt32

	}
	flag := 1
	if dividend < 0 {
		flag = -flag
		dividend = -dividend
	}
	if divisor < 0 {
		flag = -flag
		divisor = -divisor
	}
	var res int
	for dividend >= divisor {
		temp := divisor
		p := 1
		for dividend >= temp<<1 {
			temp <<= 1
			p <<= 1
		}
		res += p
		dividend -= temp
	}

	return res * flag

}
