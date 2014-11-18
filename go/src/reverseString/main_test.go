package main

import "testing"

func TestReverse(t *testing.T) {
	if Reverse("abcdef") != "fedcba" {
		t.Log("反转字符串错误")
		t.Fail()
	}
}
