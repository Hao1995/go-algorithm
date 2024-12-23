package encodeanddecodestrings

import (
	"reflect"
	"testing"
)

func TestEncode(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected string
	}{
		{
			name:     "Empty list",
			input:    []string{},
			expected: "",
		},
		{
			name:     "Multiple strings",
			input:    []string{"neet", "code", "love", "you"},
			expected: "4#neet4#code4#love3#you",
		},
		{
			name:     "Special characters (ASCII)",
			input:    []string{"we", "say", ":", "yes"},
			expected: "2#we3#say1#:3#yes",
		},
		{
			name:  "Mixed with empty strings and emoji",
			input: []string{"hello", "", " ", "😊", "go-lang"},
			// Explanation:
			// "hello"   => length 5   -> "5#hello"
			// ""        => length 0   -> "0#"
			// " "       => length 1   -> "1# "
			// "😊"      => length 1   -> "1#😊" (1 Unicode code point, though multiple bytes in UTF-8)
			// "go-lang" => length 7   -> "7#go-lang"
			expected: "5#hello0#1# 1#😊7#go-lang",
		},
		{
			name: "UTF-8 (Chinese, Arabic, Russian)",
			// Testing a variety of UTF-8 scripts
			input: []string{
				"你好",     // Chinese
				"مرحبا",  // Arabic
				"Привет", // Russian
				"こんにちは",  // Japanese
			},
			// Explanation of lengths:
			// "你好"        => 2 runes   -> "2#你好"
			// "مرحبا"      => 5 runes   -> "5#مرحبا"
			// "Привет"     => 6 runes   -> "6#Привет"
			// "こんにちは" => 5 runes   -> "5#こんにちは"
			expected: "2#你好5#مرحبا6#Привет5#こんにちは",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := encode(tt.input)
			if result != tt.expected {
				t.Errorf("encode(%v) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestDecode(t *testing.T) {
	tests := []struct {
		name     string
		encoded  string
		expected []string
	}{
		{
			name:     "Empty string",
			encoded:  "",
			expected: []string{},
		},
		{
			name:     "Multiple strings",
			encoded:  "4#neet4#code4#love3#you",
			expected: []string{"neet", "code", "love", "you"},
		},
		{
			name:     "Special characters (ASCII)",
			encoded:  "2#we3#say1#:3#yes",
			expected: []string{"we", "say", ":", "yes"},
		},
		{
			name:     "Mixed with empty strings and emoji",
			encoded:  "5#hello0#1# 1#😊7#go-lang",
			expected: []string{"hello", "", " ", "😊", "go-lang"},
		},
		{
			name:     "UTF-8 (Chinese, Arabic, Russian)",
			encoded:  "2#你好5#مرحبا6#Привет5#こんにちは",
			expected: []string{"你好", "مرحبا", "Привет", "こんにちは"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := decode(tt.encoded)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("decode(%q) = %v; want %v", tt.encoded, result, tt.expected)
			}
		})
	}
}
