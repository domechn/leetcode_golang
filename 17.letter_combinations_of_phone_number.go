// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package main

var phoneMap = map[string][]string{
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	return letterCombinationsDFS(digits[:len(digits)-1], phoneMap[string(digits[len(digits)-1])])
}

func letterCombinationsDFS(digits string, res []string) []string {
	if len(digits) == 0 {
		return res
	}
	flag := digits[len(digits)-1]
	digits = digits[:len(digits)-1]

	flagStrs := phoneMap[string(flag)]

	var temp []string

	for i := 0; i < len(flagStrs); i++ {
		for j := 0; j < len(res); j++ {
			temp = append(temp, string(flagStrs[i])+res[j])
		}
	}
	res = temp
	return letterCombinationsDFS(digits, res)
}
