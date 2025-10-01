package validations_test

import (
	"testing"

	"github.com/kishankumarhs/fnkit/validations"
)

func TestIsEmail(t *testing.T) {
	if !validations.IsEmail("foo@bar.com") {
		t.Error("expected valid email")
	}
	if validations.IsEmail("foo@bar") {
		t.Error("expected invalid email")
	}
}

func TestIsURL(t *testing.T) {
	if !validations.IsURL("https://example.com") {
		t.Error("expected valid url")
	}
	if validations.IsURL("not a url") {
		t.Error("expected invalid url")
	}
}

func TestIsUUID(t *testing.T) {
	if !validations.IsUUID("123e4567-e89b-12d3-a456-426614174000") {
		t.Error("expected valid uuid")
	}
	if validations.IsUUID("not-a-uuid") {
		t.Error("expected invalid uuid")
	}
}

func TestIsAlpha(t *testing.T) {
	if !validations.IsAlpha("abcXYZ") {
		t.Error("IsAlpha failed for letters")
	}
	if validations.IsAlpha("abc123") {
		t.Error("IsAlpha failed for alnum")
	}
}

func TestIsNumeric(t *testing.T) {
	if !validations.IsNumeric("12345") {
		t.Error("IsNumeric failed for digits")
	}
	if validations.IsNumeric("12a45") {
		t.Error("IsNumeric failed for nondigits")
	}
}

func TestIsAlnum(t *testing.T) {
	if !validations.IsAlnum("abc123") {
		t.Error("IsAlnum failed for alnum")
	}
	if validations.IsAlnum("abc-123") {
		t.Error("IsAlnum failed for non-alnum")
	}
}

func TestIsHex(t *testing.T) {
	if !validations.IsHex("deadBEEF") {
		t.Error("IsHex failed for hex")
	}
	if validations.IsHex("xyz") {
		t.Error("IsHex failed for non-hex")
	}
}

func TestIsIP(t *testing.T) {
	if !validations.IsIP("127.0.0.1") {
		t.Error("IsIP failed for IPv4")
	}
	if !validations.IsIP("::1") {
		t.Error("IsIP failed for IPv6")
	}
	if validations.IsIP("not.an.ip") {
		t.Error("IsIP failed for invalid")
	}
}

func TestIsLower(t *testing.T) {
	if !validations.IsLower("abc") {
		t.Error("IsLower failed for lower")
	}
	if validations.IsLower("aBc") {
		t.Error("IsLower failed for mixed")
	}
}

func TestIsUpper(t *testing.T) {
	if !validations.IsUpper("ABC") {
		t.Error("IsUpper failed for upper")
	}
	if validations.IsUpper("AbC") {
		t.Error("IsUpper failed for mixed")
	}
}

func TestIsASCII(t *testing.T) {
	if !validations.IsASCII("abc123") {
		t.Error("IsASCII failed for ascii")
	}
	if validations.IsASCII("a😊b") {
		t.Error("IsASCII failed for non-ascii")
	}
}

func TestIsPrintable(t *testing.T) {
	if !validations.IsPrintable("abc 123") {
		t.Error("IsPrintable failed for printable")
	}
	if validations.IsPrintable("abc\x00") {
		t.Error("IsPrintable failed for non-printable")
	}
}

func TestIsPhone(t *testing.T) {
	if !validations.IsPhone("+1-800-555-1234") {
		t.Error("IsPhone failed for valid")
	}
	if validations.IsPhone("notaphone") {
		t.Error("IsPhone failed for invalid")
	}
}

func TestIsCreditCard(t *testing.T) {
	if !validations.IsCreditCard("4111 1111 1111 1111") {
		t.Error("IsCreditCard failed for valid")
	}
	if validations.IsCreditCard("1234 5678 9012 3456") {
		t.Error("IsCreditCard failed for invalid")
	}
}

func TestToString(t *testing.T) {
	if validations.ToString(123) != "123" {
		t.Error("ToString failed")
	}
}

func TestToInt(t *testing.T) {
	if n, ok := validations.ToInt("42"); !ok || n != 42 {
		t.Error("ToInt string failed")
	}
	if n, ok := validations.ToInt(3.0); !ok || n != 3 {
		t.Error("ToInt float failed")
	}
}

func TestToFloat64(t *testing.T) {
	if f, ok := validations.ToFloat64("3.14"); !ok || f != 3.14 {
		t.Error("ToFloat64 string failed")
	}
	if f, ok := validations.ToFloat64(2); !ok || f != 2.0 {
		t.Error("ToFloat64 int failed")
	}
}
