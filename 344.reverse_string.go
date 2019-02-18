/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 344.reverse_string.go
#   Created       : 2019-02-13 10:54:36
#   Last Modified : 2019-02-13 10:54:36
#   Describe      :
#
# ====================================================*/
package leet

func reverseString(s []byte) {
	begin := 0
	end := len(s) - 1
	for begin < end {
		s[begin], s[end] = s[end], s[begin]
		begin++
		end--
	}
}
