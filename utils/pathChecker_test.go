package utils

import (
	"testing"
)

func TestIsBase64(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"aGVsbG8gd29ybGQ=", true},
		{"notbase64", false},
		{"", false},
		{"aGVsbG8", false}, // not valid base64, as it's not a multiple of 4
	}

	for _, test := range tests {
		result := isBase64(test.input)
		if result != test.expected {
			t.Errorf("isBase64(%s) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsLocalFile(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"C:\\Windows\\System32\\cmd.exe", true},
		{"C:\\nonexistent\\file.txt", false},
		{"", false},
	}

	for _, test := range tests {
		result := isLocalFile(test.input)
		if result != test.expected {
			t.Errorf("isLocalFile(%s) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestIsURL(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"https://www.example.com", true},
		{"ftp://ftp.example.com", true},
		{"not a url", false},
		{"", false},
	}

	for _, test := range tests {
		result := isURL(test.input)
		if result != test.expected {
			t.Errorf("isURL(%s) = %v; want %v", test.input, result, test.expected)
		}
	}
}

func TestCheckPath(t *testing.T) {
	tests := []struct {
		input    string
		expected uint
	}{
		{"aGVsbG8gd29ybGQ=", base64Data},
		{"C:\\Windows\\System32\\cmd.exe", localPath},
		{"https://www.example.com", netPath},
		{"not a valid path", unknown},
	}

	for _, test := range tests {
		result := CheckPath(test.input)
		if result != test.expected {
			t.Errorf("CheckPath(%s) = %v; want %v", test.input, result, test.expected)
		}
	}
}
