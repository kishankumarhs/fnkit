package fnkit_test

import (
	"fnkit"
	"testing"
	"unicode"
)

func TestRepeat(t *testing.T) {
	if got := fnkit.Repeat("ab", 3); got != "ababab" {
		t.Errorf("Repeat() = %q, want %q", got, "ababab")
	}
}

func TestReplaceAll(t *testing.T) {
	if got := fnkit.ReplaceAll("foo bar foo", "foo", "baz"); got != "baz bar baz" {
		t.Errorf("ReplaceAll() = %q", got)
	}
}

func TestHasPrefixSuffix(t *testing.T) {
	s := "hello world"
	if !fnkit.HasPrefix(s, "hello") || !fnkit.HasSuffix(s, "world") {
		t.Errorf("HasPrefix/HasSuffix failed")
	}
}

func TestContainsCount(t *testing.T) {
	s := "banana"
	if !fnkit.Contains(s, "nan") {
		t.Errorf("Contains() failed")
	}
	if fnkit.Count(s, "na") != 2 {
		t.Errorf("Count() = %d, want 2", fnkit.Count(s, "na"))
	}
}

func TestReverseString(t *testing.T) {
	if got := fnkit.ReverseString("abc"); got != "cba" {
		t.Errorf("ReverseString() = %q", got)
	}
	if got := fnkit.ReverseString("a😊b"); got != "b😊a" {
		t.Errorf("ReverseString() unicode = %q", got)
	}
}

func TestIsAlphaNumeric(t *testing.T) {
	if !fnkit.IsAlpha("abcXYZ") || fnkit.IsAlpha("abc123") {
		t.Errorf("IsAlpha failed")
	}
	if !fnkit.IsNumeric("12345") || fnkit.IsNumeric("12a45") {
		t.Errorf("IsNumeric failed")
	}
}

func TestCapitalize(t *testing.T) {
	if got := fnkit.Capitalize("hELLO"); got != "Hello" {
		t.Errorf("Capitalize() = %q", got)
	}
}

func TestStripLeftRight(t *testing.T) {
	s := "  hello  "
	if fnkit.StripLeft(s) != "hello  " || fnkit.StripRight(s) != "  hello" {
		t.Errorf("StripLeft/StripRight failed")
	}
}

func TestPartition(t *testing.T) {
	a, sep, b := fnkit.Partition("foo-bar-baz", "-")
	if a != "foo" || sep != "-" || b != "bar-baz" {
		t.Errorf("Partition() = %q, %q, %q", a, sep, b)
	}
	a, sep, b = fnkit.Rpartition("foo-bar-baz", "-")
	if a != "foo-bar" || sep != "-" || b != "baz" {
		t.Errorf("Rpartition() = %q, %q, %q", a, sep, b)
	}
}

func TestWords(t *testing.T) {
	s := "Go is  awesome"
	words := fnkit.Words(s)
	if len(words) != 3 || words[0] != "Go" || words[2] != "awesome" {
		t.Errorf("Words() = %v", words)
	}
}

func TestCaseConversions(t *testing.T) {
	s := "hello world_test-case"
	if fnkit.CamelCase(s) != "helloWorldTestCase" {
		t.Errorf("CamelCase() = %q", fnkit.CamelCase(s))
	}
	if fnkit.SnakeCase(s) != "hello_world_test_case" {
		t.Errorf("SnakeCase() = %q", fnkit.SnakeCase(s))
	}
	if fnkit.KebabCase(s) != "hello-world-test-case" {
		t.Errorf("KebabCase() = %q", fnkit.KebabCase(s))
	}
}

func TestIsUpperLower(t *testing.T) {
	if !fnkit.IsUpper("ABC") || fnkit.IsUpper("AbC") {
		t.Errorf("IsUpper failed")
	}
	if !fnkit.IsLower("abc") || fnkit.IsLower("aBc") {
		t.Errorf("IsLower failed")
	}
}

func TestSwapCase(t *testing.T) {
	if got := fnkit.SwapCase("aBc"); got != "AbC" {
		t.Errorf("SwapCase() = %q", got)
	}
}

func TestRemoveKeep(t *testing.T) {
	s := "a1b2c3"
	onlyDigits := fnkit.Keep(s, unicode.IsDigit)
	if onlyDigits != "123" {
		t.Errorf("Keep() = %q", onlyDigits)
	}
	noDigits := fnkit.Remove(s, unicode.IsDigit)
	if noDigits != "abc" {
		t.Errorf("Remove() = %q", noDigits)
	}
}

func TestPadCenter(t *testing.T) {
	if got := fnkit.PadCenter("hi", 6, '*'); got != "**hi**" {
		t.Errorf("PadCenter() = %q", got)
	}
}
