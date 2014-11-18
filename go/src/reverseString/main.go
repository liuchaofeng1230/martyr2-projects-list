package main

import (
	"fmt"
	"utils"
)

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {

	orgText := utils.StdIn("请输入待反转的字符串:")
	s := Reverse(orgText)

	fmt.Printf("反转后的字符串: %s\n", string(s))
}
