package utils

import "testing"

func TestReverseString(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "Empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "Single character",
			input:    "a",
			expected: "a",
		},
		{
			name:     "Simple English word",
			input:    "hello",
			expected: "olleh",
		},
		{
			name:     "English sentence with spaces",
			input:    "hello world",
			expected: "dlrow olleh",
		},
		{
			name:     "Russian word",
			input:    "привет",
			expected: "тевирп",
		},
		{
			name:     "Russian sentence with spaces",
			input:    "привет мир",
			expected: "рим тевирп",
		},
		{
			name:     "Mixed languages",
			input:    "hello привет",
			expected: "тевирп olleh",
		},
		{
			name:     "Special characters",
			input:    "hello! привет!",
			expected: "!тевирп !olleh",
		},
		{
			name:     "Numbers",
			input:    "123 456",
			expected: "654 321",
		},
		{
			name:     "Emojis",
			input:    "👋 hello 🌍",
			expected: "🌍 olleh 👋",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ReverseString(tt.input)
			if result != tt.expected {
				t.Errorf("ReverseString(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
} 