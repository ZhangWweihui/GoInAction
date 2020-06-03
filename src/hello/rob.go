/**
leetcode 打家劫舍 https://leetcode-cn.com/problems/house-robber/submissions/
*/
package main

import "fmt"

func main() {
	var arr = []int{2, 4, 2}
	fmt.Println(rob2(arr))
}

func rob(nums []int) int {
	even, odd := 0, 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			even += nums[i]
		} else {
			odd += nums[i]
		}
	}
	if even > odd {
		return even
	} else {
		return odd
	}
}

func rob2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	}
	money := 0
	distance := len(nums) - 1
	num1, num2 := 0, 0
	for k := 2; k <= distance; k++ {
		num1 += nums[0]
		num2 += nums[1]
		for i := 0; i < len(nums) && (i+1)*k < len(nums); i++ {
			num1 += nums[(i+1)*k]
		}
		for i := 0; i < len(nums) && (i+1)*k+1 < len(nums); i++ {
			num2 += nums[(i+1)*k+1]
		}
		if num1 >= num2 && num1 > money {
			money = num1
		} else if num2 >= num1 && num2 > money {
			money = num2
		}
		num1, num2 = 0, 0
	}

	return money
}
