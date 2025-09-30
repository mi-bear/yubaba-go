package main

import (
	"testing"
)

// 日本語文字列からランダムに1文字取ります。
func TestStealName(t *testing.T) {
	name := "千尋"
	got := stealName(name)

	if got != "千" && got != "尋" {
		t.Errorf("unexpected character: got=%s", got)
	}
}

func TestStealNameSingleChar(t *testing.T) {
	name := "湯"
	got := stealName(name)
	if got != "湯" {
		t.Errorf("expected '湯', got=%s", got)
	}
}
