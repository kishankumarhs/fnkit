package fnkit

import (
	"strings"
	"unicode"
)

// Repeat returns a new string consisting of count copies of s.
func Repeat(s string, count int) string {
	return strings.Repeat(s, count)
}

// ReplaceAll replaces all occurrences of old with new in s.
func ReplaceAll(s, old, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// HasPrefix returns true if s starts with prefix.
func HasPrefix(s, prefix string) bool {
	return strings.HasPrefix(s, prefix)
}

// HasSuffix returns true if s ends with suffix.
func HasSuffix(s, suffix string) bool {
	return strings.HasSuffix(s, suffix)
}

// Contains returns true if s contains substr.
func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// Count returns the number of non-overlapping instances of substr in s.
func Count(s, substr string) int {
	return strings.Count(s, substr)
}

// Reverse returns the string s reversed (Unicode-safe).
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// IsAlpha returns true if all characters in s are letters.
func IsAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return len(s) > 0
}

// IsNumeric returns true if all characters in s are digits.
func IsNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return len(s) > 0
}

// Capitalize returns a copy of s with the first letter capitalized.
func Capitalize(s string) string {
	if s == "" {
		return s
	}
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}

// StripLeft removes all trailing whitespace from s.
func StripLeft(s string) string {
	return strings.TrimLeftFunc(s, unicode.IsSpace)
}

// StripRight removes all trailing whitespace from s.
func StripRight(s string) string {
	return strings.TrimRightFunc(s, unicode.IsSpace)
}

// Partition splits s at the first occurrence of sep. Returns before, sep, after.
func Partition(s, sep string) (string, string, string) {
	idx := strings.Index(s, sep)
	if idx == -1 {
		return s, "", ""
	}
	return s[:idx], sep, s[idx+len(sep):]
}

// Rpartition splits s at the last occurrence of sep. Returns before, sep, after.
func Rpartition(s, sep string) (string, string, string) {
	idx := strings.LastIndex(s, sep)
	if idx == -1 {
		return s, "", ""
	}
	return s[:idx], sep, s[idx+len(sep):]
}

// Words splits s into fields separated by whitespace.
func Words(s string) []string {
	return strings.Fields(s)
}

// CamelCase converts a string to camelCase.
func CamelCase(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}
	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
	for i := range words {
		if i == 0 {
			words[i] = strings.ToLower(words[i])
		} else {
			words[i] = Capitalize(words[i])
		}
	}
	return strings.Join(words, "")
}

// SnakeCase converts a string to snake_case.
func SnakeCase(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}
	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return strings.Join(words, "_")
}

// KebabCase converts a string to kebab-case.
func KebabCase(s string) string {
	s = strings.TrimSpace(s)
	if s == "" {
		return s
	}
	words := strings.FieldsFunc(s, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsDigit(r)
	})
	for i := range words {
		words[i] = strings.ToLower(words[i])
	}
	return strings.Join(words, "-")
}

// IsUpper returns true if all letters in s are uppercase.
func IsUpper(s string) bool {
	has := false
	for _, r := range s {
		if unicode.IsLetter(r) {
			has = true
			if !unicode.IsUpper(r) {
				return false
			}
		}
	}
	return has
}

// IsLower returns true if all letters in s are lowercase.
func IsLower(s string) bool {
	has := false
	for _, r := range s {
		if unicode.IsLetter(r) {
			has = true
			if !unicode.IsLower(r) {
				return false
			}
		}
	}
	return has
}

// SwapCase returns a copy of s with upper and lower case letters swapped.
func SwapCase(s string) string {
	runes := []rune(s)
	for i, r := range runes {
		if unicode.IsUpper(r) {
			runes[i] = unicode.ToLower(r)
		} else if unicode.IsLower(r) {
			runes[i] = unicode.ToUpper(r)
		}
	}
	return string(runes)
}

// Remove returns a copy of s with all runes for which remove(r) is true removed.
func Remove(s string, remove func(rune) bool) string {
	runes := []rune(s)
	out := runes[:0]
	for _, r := range runes {
		if !remove(r) {
			out = append(out, r)
		}
	}
	return string(out)
}

// Keep returns a copy of s with only runes for which keep(r) is true.
func Keep(s string, keep func(rune) bool) string {
	runes := []rune(s)
	out := runes[:0]
	for _, r := range runes {
		if keep(r) {
			out = append(out, r)
		}
	}
	return string(out)
}

// PadCenter pads s on both sides with padChar to center it in a string of length width.
func PadCenter(s string, width int, padChar rune) string {
	if len(s) >= width {
		return s
	}
	pad := width - len(s)
	left := pad / 2
	right := pad - left
	return strings.Repeat(string(padChar), left) + s + strings.Repeat(string(padChar), right)
}
