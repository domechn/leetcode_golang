/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 121.best_time_to_buy_and_sell_stock.go
#   Created       : 2019-03-13 13:43:15
#   Last Modified : 2019-03-13 13:43:15
#   Describe      :
#
# ====================================================*/
package main

func maxProfit(prices []int) int {
	var max, t int

	for i := 0; i < len(prices)-1; i++ {
		temp := prices[i+1] - prices[i]
		if t+temp > 0 {
			t += temp
		} else {
			t = 0
		}
		if max < t {
			max = t
		}
	}
	return max
}
