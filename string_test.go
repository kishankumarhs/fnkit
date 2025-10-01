package fnkit

import (
	"testing"
	"unicode"
)

func TestRepeat(t *testing.T) {
	if got := Repeat("ab", 3); got != "ababab" {
		t.Errorf("Repeat() = %q, want %q", got, "ababab")
	}
}

func TestReplaceAll(t *testing.T) {
	if got := ReplaceAll("foo bar foo", "foo", "baz"); got != "baz bar baz" {
		t.Errorf("ReplaceAll() = %q", got)
	}
}

func TestHasPrefixSuffix(t *testing.T) {
	s := "hello world"
	if !HasPrefix(s, "hello") || !HasSuffix(s, "world") {
		t.Errorf("HasPrefix/HasSuffix failed")
	}
}

func TestContainsCount(t *testing.T) {
	s := "banana"
	if !Contains(s, "nan") {
		t.Errorf("Contains() failed")
	}
	if Count(s, "na") != 2 {
		t.Errorf("Count() = %d, want 2", Count(s, "na"))
	}
}

func TestReverseString(t *testing.T) {
	if got := ReverseString("abc"); got != "cba" {
		t.Errorf("ReverseString() = %q", got)
	}
	if got := ReverseString("a😊b"); got != "b😊a" {
		t.Errorf("ReverseString() unicode = %q", got)
	}
}

func TestIsAlphaNumeric(t *testing.T) {
	if !IsAlpha("abcXYZ") || IsAlpha("abc123") {
		t.Errorf("IsAlpha failed")
	}
	if !IsNumeric("12345") || IsNumeric("12a45") {
		t.Errorf("IsNumeric failed")
	}
}

func TestCapitalize(t *testing.T) {
	if got := Capitalize("hELLO"); got != "Hello" {
		t.Errorf("Capitalize() = %q", got)
	}
}

func TestStripLeftRight(t *testing.T) {
	s := "  hello  "
	if StripLeft(s) != "hello  " || StripRight(s) != "  hello" {
		t.Errorf("StripLeft/StripRight failed")
	}
}

func TestPartition(t *testing.T) {
	a, sep, b := Partition("foo-bar-baz", "-")
	if a != "foo" || sep != "-" || b != "bar-baz" {
		t.Errorf("Partition() = %q, %q, %q", a, sep, b)
	}
	a, sep, b = Rpartition("foo-bar-baz", "-")
	if a != "foo-bar" || sep != "-" || b != "baz" {
		t.Errorf("Rpartition() = %q, %q, %q", a, sep, b)
	}
}

func TestWords(t *testing.T) {
	s := "Go is  awesome"
	words := Words(s)
	if len(words) != 3 || words[0] != "Go" || words[2] != "awesome" {
		t.Errorf("Words() = %v", words)
	}
}

func TestCaseConversions(t *testing.T) {
	s := "hello world_test-case"
	if CamelCase(s) != "helloWorldTestCase" {
		t.Errorf("CamelCase() = %q", CamelCase(s))
	}
	if SnakeCase(s) != "hello_world_test_case" {
		t.Errorf("SnakeCase() = %q", SnakeCase(s))
	}
	if KebabCase(s) != "hello-world-test-case" {
		t.Errorf("KebabCase() = %q", KebabCase(s))
	}
}

func TestIsUpperLower(t *testing.T) {
	if !IsUpper("ABC") || IsUpper("AbC") {
		t.Errorf("IsUpper failed")
	}
	if !IsLower("abc") || IsLower("aBc") {
		t.Errorf("IsLower failed")
	}
}

func TestSwapCase(t *testing.T) {
	if got := SwapCase("aBc"); got != "AbC" {
		t.Errorf("SwapCase() = %q", got)
	}
}

func TestRemoveKeep(t *testing.T) {
	s := "a1b2c3"
	onlyDigits := Keep(s, unicode.IsDigit)
	if onlyDigits != "123" {
		t.Errorf("Keep() = %q", onlyDigits)
	}
	noDigits := Remove(s, unicode.IsDigit)
	if noDigits != "abc" {
		t.Errorf("Remove() = %q", noDigits)
	}
}

func TestPadCenter(t *testing.T) {
	if got := PadCenter("hi", 6, '*'); got != "**hi**" {
		t.Errorf("PadCenter() = %q", got)
	}
}
