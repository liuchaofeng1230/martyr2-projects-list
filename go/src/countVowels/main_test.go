package main

import "testing"

func TestCountVowels(t *testing.T) {
	store := CountVowels("aabbiiru")

	if store["a"] != 2 || store["e"] != 0 || store["i"] != 2 || store["o"] != 0 || store["u"] != 1 {
		t.Log("统计元音字母数量错误")
		t.Fail()
	}
}
