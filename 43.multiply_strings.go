package main

import "fmt"

func main() {
	fmt.Println(multiply("123", "24"))
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	n1 := len(num1) - 1
	n2 := len(num2) - 1

	res := make([]int, n1+n2+2)

	for i := 0; i < n1+1; i++ {
		for j := 0; j < n2+1; j++ {
			temp := (int(num1[n1-i] - '0')) * (int(num2[n2-j] - '0'))
			res[len(res)-i-j-1] += temp % 10
			res[len(res)-i-j-2] += temp / 10
		}
	}

	for i := len(res) - 1; i > 0; i-- {
		if res[i] > 9 {
			temp := res[i] / 10
			res[i] %= 10
			res[i-1] += temp
		}
	}
	var resStr string

	flag := false
	for _, v := range res {
		if v != 0 || flag {
			flag = true
			resStr += string(v + '0')
		}
	}
	return resStr
}

