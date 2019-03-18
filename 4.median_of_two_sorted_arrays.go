/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 4.median_of_two_sorted_arrays.go
#   Created       : 2019-02-19 18:01:38
#   Last Modified : 2019-02-19 18:01:38
#   Describe      :
#
# ====================================================*/
package main

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	return float64(robbin(nums1, 0, nums2, 0, (m+n+1)/2)+robbin(nums1, 0, nums2, 0, (m+n+2)/2)) / 2.0
}

func robbin(nums1 []int, i int, nums2 []int, j int, k int) int {
	if i >= len(nums1) {
		return nums2[j+k-1]
	}
	if j >= len(nums2) {
		return nums1[i+k-1]
	}
	if k == 1 {
		return min(nums1[i], nums2[j])
	}

	var mid1, mid2 = math.MaxInt32, math.MaxInt32
	if i+k/2-1 < len(nums1) {
		mid1 = nums1[i+k/2-1]
	}
	if j+k/2-1 < len(nums2) {
		mid2 = nums2[j+k/2-1]
	}
	if mid1 < mid2 {
		return robbin(nums1, i+k/2, nums2, j, k-k/2)
	}
	return robbin(nums1, i, nums2, j+k/2, k-k/2)

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
