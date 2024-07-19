package utils

import (
	"os/user"
	"reflect"
	"sync"
	"testing"
)

func TestGetHomeDir(t *testing.T) {
	var once sync.Once
	onceHomeDir = once // 重置 sync.Once 以确保每次测试运行时都执行一次 Do 函数

	// 获取当前用户的 home directory
	u, err := user.Current()
	if err != nil {
		t.Fatalf("Failed to get current user: %v", err)
	}

	expected := u.HomeDir
	actual := GetHomeDir()

	if actual != expected {
		t.Errorf("Expected home dir: %v, but got: %v", expected, actual)
	}
}

func TestStrimList(t *testing.T) {
	tests := []struct {
		input  []string
		output []string
	}{
		{input: nil, output: nil},
		{input: []string{}, output: []string{}},
		{input: []string{"a", "b", "c"}, output: []string{"a", "b", "c"}},
		{input: []string{"a", "", "c"}, output: []string{"a", "c"}},
		{input: []string{"", "", ""}, output: []string{}},
		{input: []string{"a", "", "b", "", "c", ""}, output: []string{"a", "b", "c"}},
	}

	for _, test := range tests {
		result := StrimList(test.input)
		if !reflect.DeepEqual(result, test.output) {
			t.Errorf("StrimList(%v) = %v; want %v", test.input, result, test.output)
		}
	}
}

func TestDeleteAfterLastCharacter(t *testing.T) {
	tests := []struct {
		input    string
		char     string
		expected string
	}{
		{"hello world", " ", "hello"},
		{"hello world", "o", "hello w"},
		{"hello world", "d", "hello worl"},
		{"hello world", "x", "hello world"},
		{"", "x", ""},
		{"hello", "", "hello"},
	}

	for _, test := range tests {
		result := DeleteAfterLastCharacter(test.input, test.char)
		if result != test.expected {
			t.Errorf("DeleteAfterLastCharacter(%q, %q) = %q; want %q", test.input, test.char, result, test.expected)
		}
	}
}
