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
	for i, l := 0, int(math.Floor(float64(length/2))); i < l; i++ {
		s[length-i-1] = s[i]
	}

	result = string(s) == str
	return
}

func main() {
	input := utils.StdIn("请输入字符串")
	result := isPalindrome(input)
	fmt.Printf("是否回文: %s", strconv.FormatBool(result)+"\n")
}
