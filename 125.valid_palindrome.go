/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 125.valid_palindrome.go
#   Created       : 2019-03-04 17:44:17
#   Last Modified : 2019-03-04 17:44:17
#   Describe      :
#
# ====================================================*/
package leet

func isPalindrome(s string) bool {
	begin := 0
	end := len(s) - 1

	for begin < end {
		if (s[begin] >= 'a' && s[begin] <= 'z') || (s[begin] >= '0' && s[begin] <= '9') || (s[begin] >= 'A' && s[begin] <= 'Z') {

		} else {
			begin++
			continue
		}

		if (s[end] >= 'a' && s[end] <= 'z') || (s[end] >= '0' && s[end] <= '9') || (s[end] >= 'A' && s[end] <= 'Z') {

		} else {
			end--
			continue
		}

		if s[begin] != s[end] {
			if abs(int(s[begin])-int(s[end])) != 32 {
				return false
			}
			if (s[end] >= '0' && s[end] <= '9') || (s[begin] >= '0' && s[begin] <= '9') {
				return false
			}
		}
		begin++
		end--
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
