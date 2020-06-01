package main

import "fmt"

func main() {
	fmt.Println(findRepeatedDnaSequences("AAAAAAAAAAA"))
}

func findRepeatedDnaSequences(s string) []string {
	if len(s) < 10 {
		return []string{}
	}
	var temp = make(map[string]int, len(s)-10)
	for i := 0; i+9 < len(s); i++ {
		temp[s[i:i+10]]++
	}

	var res = []string{}
	for k, v := range temp {
		if v > 1 {
			res = append(res, k)
		}
	}
	return res
}
