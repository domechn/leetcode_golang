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
// class Solution {
// 	public:
// 		string countAndSay(int n) {
// 			if (n <= 0) return "";
// 			string res = "1";
// 			while (--n) {
// 				string cur = "";
// 				for (int i = 0; i < res.size(); ++i) {
// 					int cnt = 1;
// 					while (i + 1 < res.size() && res[i] == res[i + 1]) {
// 						++cnt;
// 						++i;
// 					}
// 					cur += to_string(cnt) + res[i];
// 				}
// 				res = cur;
// 			}
// 			return res;
// 		}
// 	};
func countAndSay(n int) string {
	if n <= 0 {
		return ""
	}
	res := "1"
	for n > 0 {
		n --

	}
}
