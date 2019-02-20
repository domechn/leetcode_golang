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

func main() {

}

func permute(nums []int) [][]int {
	res := make([][]int{},500)
	for i	:= 0; i	< len(nums); i	+ {
		permuteDFS(res,0,ma)
	}
	
}

func permuteDFS(res [][]int,level int , visited ,nums,out []int) {
	if level == len(nums) {
		 res = append( res, out)
		return 
	}

for i := 0; i < len(nums) && visited[i] == 0; i++ {
	visited[i]= 1
	 out= append( out,nums[i])
	 permuteDFS(res,level+1,visited,nums,out)
}

}
