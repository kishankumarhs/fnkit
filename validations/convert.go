package validations

import (
	"fmt"
	"reflect"
	"strconv"
)

// ToString converts any value to a string.
func ToString(v any) string {
	return fmt.Sprint(v)
}

// ToInt tries to convert any value to an int. Returns (value, true) if successful.
func ToInt(v any) (int, bool) {
	switch x := v.(type) {
	case int:
		return x, true
	case int8, int16, int32, int64:
		return int(reflect.ValueOf(x).Int()), true
	case uint, uint8, uint16, uint32, uint64:
		return int(reflect.ValueOf(x).Uint()), true
	case float32, float64:
		return int(reflect.ValueOf(x).Float()), true
	case string:
		n, err := strconv.Atoi(x)
		return n, err == nil
	default:
		return 0, false
	}
}

// ToFloat64 tries to convert any value to a float64. Returns (value, true) if successful.
func ToFloat64(v any) (float64, bool) {
	switch x := v.(type) {
	case float64:
		return x, true
	case float32:
		return float64(x), true
	case int, int8, int16, int32, int64:
		return float64(reflect.ValueOf(x).Int()), true
	case uint, uint8, uint16, uint32, uint64:
		return float64(reflect.ValueOf(x).Uint()), true
	case string:
		f, err := strconv.ParseFloat(x, 64)
		return f, err == nil
	default:
		return 0, false
	}
}
