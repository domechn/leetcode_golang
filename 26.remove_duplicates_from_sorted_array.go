// Copyright (c) 2018, dmc (814172254@qq.com),
//
// Authors: dmc,
//
// Distribution:.
package main

import (
	"fmt"
)

func main() {
	fmt.Println(removeDuplicates([]int{0, 0, 1}))
}

func removeDuplicates(nums []int) int {
	var pre, cur int
	for cur < len(nums) {
		if nums[cur] == nums[pre] {
			cur++
		} else {
			pre++
			nums[pre] = nums[cur]
			cur++
		}
	}

	return pre + 1
}
