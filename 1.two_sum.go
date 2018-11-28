package main

type Interval struct {
	Start, End int
}

func twoSum(nums []int, target int) []int {
	res := make([]int, 2)

	temp := make(map[int]int, len(nums))

	for i := 0; i < len(nums); i++ {
		v, ok := temp[target-nums[i]]
		if ok {
			if i < v {
				res[0] = i
				res[1] = v
			} else {
				res[0] = v
				res[1] = i
			}
			break
		}
		temp[nums[i]] = i
	}
	return res
}
