/* ====================================================
#   Copyright (C)2019 All rights reserved.
#
#   Author        : domchan
#   Email         : 814172254@qq.com
#   File Name     : 88.merge_sorted_array.go
#   Created       : 2019-03-04 17:09:04
#   Last Modified : 2019-03-04 17:09:04
#   Describe      :
#
# ====================================================*/
package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	var p = m + n - 1
	m--
	n--
	for m >= 0 && n >= 0 {
		var t int
		if nums1[m] >= nums2[n] {
			t = nums1[m]
			m--
		} else {
			t = nums2[n]
			n--
		}

		nums1[p] = t
		p--
	}

	for n >= 0 {
		nums1[p] = nums2[n]
		n--
		p--
	}

}
