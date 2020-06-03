/**
leetcode IP地址无效化 https://leetcode-cn.com/problems/defanging-an-ip-address/
*/
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println(defangIPaddr("125.125.125.125"))

	str := "扶摇直上九万里"
	fmt.Println("str length : ", len(str))
	fmt.Println("str rune length : ", utf8.RuneCountInString(str))
	fmt.Println("str rune length2 : ", len([]rune(str)))

	for k, v := range str {
		fmt.Printf("v type: %T\n", v)
		fmt.Println(v, k)
	}

	arr()

	arr1()
}

func defangIPaddr(address string) string {
	//return strings.Replace(address, ".", "[.]", -1)
	var chars = []rune(address)
	var newChars []rune
	spe := []rune("[.]")
	for i := 0; i < len(chars); i++ {
		if chars[i] == spe[1] {
			newChars = append(newChars, spe[0], spe[1], spe[2])
		} else {
			newChars = append(newChars, chars[i])
		}
	}
	return string(newChars)
}

func arr() {
	var nums [10]float32
	var i, j int

	for i = 0; i < len(nums); i++ {
		nums[i] = (float32)(i+1) * 100
	}

	for j = 0; j < len(nums); j++ {
		fmt.Printf("nums[%d]=%f \n", j, nums[j])
	}
}

func arr1() {
	var nums [10]int

	for i := 0; i < 10; i++ {
		nums[i] = (i + 1) * 100
	}

	var nums1 = nums[:3]
	nums1[1] = 37
	fmt.Printf("slice nums1 len : %d, cap: %d, val:%v\n", len(nums1), cap(nums1), nums1)

	var nums2 = nums[:4]
	nums2[1] = 68
	fmt.Printf("slice nums2 len : %d, cap: %d, val:%v\n", len(nums2), cap(nums2), nums2)

	fmt.Printf("arr nums len : %d, cap: %d, val:%v\n", len(nums), cap(nums), nums)
	fmt.Printf("slice nums1 len : %d, cap: %d, val:%v\n", len(nums1), cap(nums1), nums1)
}
