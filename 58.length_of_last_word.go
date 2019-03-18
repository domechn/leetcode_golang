/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 58.length_of_last_word.go
#   Created       : 2019-02-15 19:40:23
#   Last Modified : 2019-02-15 19:40:23
#   Describe      :
#
# ====================================================*/
package main

import "strings"

func lengthOfLastWord(s string) int {
	ss := strings.Split(s, " ")

	for i := len(ss) - 1; i >= 0; i-- {
		if ss[i] != "" {
			return len(ss[i])
		}
	}

	return 0

}
