# fnkit/validations

A Go module for common data validation and type conversion utilities. Import as:

```go
import "github.com/kishankumarhs/fnkit/validations"
```

## Validators

- `IsEmail(s string) bool` — Validates email addresses.
- `IsURL(s string) bool` — Validates URLs.
- `IsUUID(s string) bool` — Validates UUIDs (v1-v5).
- `IsAlpha(s string) bool` — Letters only (Unicode).
- `IsNumeric(s string) bool` — Digits only (Unicode).
- `IsAlnum(s string) bool` — Letters or digits (Unicode).
- `IsHex(s string) bool` — Hexadecimal string.
- `IsIP(s string) bool` — IPv4 or IPv6 address.
- `IsLower(s string) bool` — All letters lowercase.
- `IsUpper(s string) bool` — All letters uppercase.
- `IsASCII(s string) bool` — All runes are ASCII.
- `IsPrintable(s string) bool` — All runes are printable.
- `IsPhone(s string) bool` — Simple phone number check.
- `IsCreditCard(s string) bool` — Credit card number (Luhn check).

## Type Conversion

- `ToString(v any) string` — Converts any value to string.
- `ToInt(v any) (int, bool)` — Converts to int if possible.
- `ToFloat64(v any) (float64, bool)` — Converts to float64 if possible.

## Example Usage

```go
import "github.com/kishankumarhs/fnkit/validations"

validations.IsEmail("foo@bar.com") // true
validations.IsUUID("123e4567-e89b-12d3-a456-426614174000") // true
validations.ToInt("42") // 42, true
```

## Tests

All functions are covered by unit tests in `validate_test.go`.

---

MIT License
