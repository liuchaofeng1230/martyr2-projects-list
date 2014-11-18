package main

import "testing"

func TestPalindrome(t *testing.T) {
	if !isPalindrome("abcdefgfedcba") || isPalindrome("abcdefgfedcb") {
		t.Log("统计元音字母数量错误")
		t.Fail()
	}
}
