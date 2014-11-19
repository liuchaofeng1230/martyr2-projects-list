package main

import (
	"fmt"
	"math"
	"strconv"
	"utils"
)

func isPalindrome(str string) (result bool) {
	s := []rune(str)
	length := len(s)
	result = true
	for i, l := 0, int(math.Floor(float64(length/2))); i < l; i++ {
		if s[length-i-1] != s[i] {
			result = false
			return
		}
	}
	return
}

func main() {
	input := utils.StdIn("请输入字符串")
	result := isPalindrome(input)
	fmt.Println("是否回文: " + strconv.FormatBool(result))
}
