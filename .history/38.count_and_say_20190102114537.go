/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 38.count_and_say.go
#   Created       : 2019-01-02 11:19:48
#   Describe      :
#
# ====================================================*/
package main

import "fmt"

func main() {
	fmt.Println(countAndSay(5))
}

func countAndSay(n int) string {
	if n <= 0 {
		return ""
	}
	res := "1"
	for n >= 0 {
		n--
		cur := ""
		for i := 0; i < len(res); i++ {
			cnt := 1
			for i+1 < len(res) && res[i] == res[i+1] {
				cnt++
				i++
			}
			cur = fmt.Sprintf("%d%s%s", string(cnt), res[i], cur)
		}
		res = cur
	}
	return res
}
