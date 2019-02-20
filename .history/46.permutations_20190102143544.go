/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 46.permutations.go
#   Created       : 2019-01-02 11:53:38
#   Describe      :
#
# ====================================================*/
package main

import "fmt"

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
}

func permute(nums []int) [][]int {
	var res [][]int
	for i := 0; i < len(nums); i++ {
		permuteDFS(&res, 0, make(map[int]int, len(nums)), nums, []int{})
	}
	return res
}

func permuteDFS(res *[][]int, level int, visited map[int]int, nums, out []int) {
	if level == len(nums)-1 {
		fmt.Println(level)
		res = append(res, out)
		return
	}
	for i := 0; i < len(nums) && visited[i] == 0; i++ {
		visited[i] = 1
		out = append(out, nums[i])
		permuteDFS(res, level+1, visited, nums, out)
		out = out[:len(out)-1]
		visited[i] = 0
	}

}
