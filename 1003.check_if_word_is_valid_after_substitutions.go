/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 1003.check_if_word_is_valid_after_substitutions.go
#   Created       : 2019-03-04 15:50:30
#   Last Modified : 2019-03-04 15:50:30
#   Describe      :
#
# ====================================================*/
package leet

func isValid(S string) bool {
	var t []int32
	for _, v := range S {
		if v == 'a' || v == 'b' {
			t = append(t, v)
		} else if v == 'c' {
			if len(t) >= 2 && t[len(t)-1] == 'b' && t[len(t)-2] == 'a' {
				t = t[:len(t)-2]
			}
		}
	}
	return len(t) == 0
}
