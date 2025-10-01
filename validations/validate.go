package validations

import (
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"unicode"
)

// IsAlpha returns true if s contains only letters (Unicode).
func IsAlpha(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return len(s) > 0
}

// IsNumeric returns true if s contains only digits (Unicode).
func IsNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return len(s) > 0
}

// IsAlnum returns true if s contains only letters or digits (Unicode).
func IsAlnum(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			return false
		}
	}
	return len(s) > 0
}

// IsHex returns true if s is a valid hexadecimal string (0-9, a-f, A-F).
func IsHex(s string) bool {
	for _, r := range s {
		if !('0' <= r && r <= '9') && !('a' <= r && r <= 'f') && !('A' <= r && r <= 'F') {
			return false
		}
	}
	return len(s) > 0
}

// IsIP returns true if s is a valid IPv4 or IPv6 address.
func IsIP(s string) bool {
	return net.ParseIP(s) != nil
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

// IsASCII returns true if all runes in s are ASCII.
func IsASCII(s string) bool {
	for _, r := range s {
		if r > 127 {
			return false
		}
	}
	return len(s) > 0
}

// IsPrintable returns true if all runes in s are printable.
func IsPrintable(s string) bool {
	for _, r := range s {
		if !unicode.IsPrint(r) {
			return false
		}
	}
	return len(s) > 0
}

// IsPhone returns true if s looks like a phone number (simple, digits, optional +, -, spaces).
func IsPhone(s string) bool {
	phoneRe := regexp.MustCompile(`^\+?[0-9\-\s]{7,}$`)
	return phoneRe.MatchString(s)
}

// IsCreditCard returns true if s looks like a credit card number (Luhn check, 13-19 digits).
func IsCreditCard(s string) bool {
	digits := regexp.MustCompile(`\D`).ReplaceAllString(s, "")
	if len(digits) < 13 || len(digits) > 19 {
		return false
	}
	var sum int
	alt := false
	for i := len(digits) - 1; i >= 0; i-- {
		n := int(digits[i] - '0')
		if alt {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
		alt = !alt
	}
	return sum%10 == 0
}

// IsEmail returns true if s is a valid email address.
func IsEmail(s string) bool {
	_, err := mail.ParseAddress(s)
	return err == nil
}

// IsURL returns true if s is a valid URL.
func IsURL(s string) bool {
	_, err := url.ParseRequestURI(s)
	return err == nil
}

var uuidRegex = regexp.MustCompile(`^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-[1-5][a-fA-F0-9]{3}-[89abAB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$`)

// IsUUID returns true if s is a valid UUID (version 1-5).
func IsUUID(s string) bool {
	return uuidRegex.MatchString(s)
}
