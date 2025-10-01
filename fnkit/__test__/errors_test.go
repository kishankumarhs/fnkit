package fnkit

import (
	"errors"
	"testing"
)

func mockDBQuery(id int) Result[string] {
	if id == 1 {
		return Ok("user1")
	}
	return Err[string](errors.New("not found"))
}

func mockNetworkCall(url string) Result[int] {
	if url == "https://ok" {
		return Ok(200)
	}
	return Err[int](errors.New("network error"))
}

func TestResultWithMockDBQuery(t *testing.T) {
	res := mockDBQuery(1)
	if !res.IsOk() || res.Value != "user1" {
		t.Errorf("Expected Ok with user1, got %+v", res)
	}
	res = mockDBQuery(2)
	if res.IsOk() || res.Err == nil || res.Value != "" {
		t.Errorf("Expected Err with zero value, got %+v", res)
	}
	if res.ValueOr("default") != "default" {
		t.Errorf("Expected ValueOr to return fallback, got %v", res.ValueOr("default"))
	}
}

func TestResultWithMockNetworkCall(t *testing.T) {
	ok := mockNetworkCall("https://ok")
	if !ok.IsOk() || ok.Value != 200 {
		t.Errorf("Expected Ok with 200, got %+v", ok)
	}
	err := mockNetworkCall("https://fail")
	if err.IsOk() || err.Err == nil || err.Value != 0 {
		t.Errorf("Expected Err with zero value, got %+v", err)
	}
	if err.ValueOr(404) != 404 {
		t.Errorf("Expected ValueOr to return fallback, got %v", err.ValueOr(404))
	}
}

func TestOkResult(t *testing.T) {
	r := Ok(42)
	if !r.IsOk() {
		t.Errorf("Expected IsOk to be true for Ok result")
	}
	if r.Value != 42 {
		t.Errorf("Expected Value to be 42, got %v", r.Value)
	}
	if r.Err != nil {
		t.Errorf("Expected Err to be nil, got %v", r.Err)
	}
	if r.ValueOr(99) != 42 {
		t.Errorf("Expected ValueOr to return 42, got %v", r.ValueOr(99))
	}
}

func TestErrResult(t *testing.T) {
	err := errors.New("fail")
	r := Err[int](err)
	if r.IsOk() {
		t.Errorf("Expected IsOk to be false for Err result")
	}
	if r.Err != err {
		t.Errorf("Expected Err to be %v, got %v", err, r.Err)
	}
	if r.Value != 0 {
		t.Errorf("Expected Value to be zero value, got %v", r.Value)
	}
	if r.ValueOr(99) != 99 {
		t.Errorf("Expected ValueOr to return fallback, got %v", r.ValueOr(99))
	}
}

func TestOkResultWithZeroValue(t *testing.T) {
	r := Ok(0)
	if !r.IsOk() {
		t.Errorf("Expected IsOk to be true for Ok result with zero value")
	}
	if r.Value != 0 {
		t.Errorf("Expected Value to be 0, got %v", r.Value)
	}
	if r.ValueOr(99) != 0 {
		t.Errorf("Expected ValueOr to return 0, got %v", r.ValueOr(99))
	}
}

func TestErrResultWithNilError(t *testing.T) {
	r := Err[string](nil)
	if !r.IsOk() {
		t.Errorf("Expected IsOk to be true for Err with nil error (should be Ok)")
	}
	if r.Err != nil {
		t.Errorf("Expected Err to be nil, got %v", r.Err)
	}
	if r.Value != "" {
		t.Errorf("Expected Value to be zero value, got %v", r.Value)
	}
}

func TestResultWithStructType(t *testing.T) {
	type S struct{ X int }
	r := Ok(S{X: 5})
	if !r.IsOk() {
		t.Errorf("Expected IsOk to be true for Ok result with struct")
	}
	if r.Value.X != 5 {
		t.Errorf("Expected Value.X to be 5, got %v", r.Value.X)
	}
	r2 := Err[S](errors.New("fail"))
	if r2.IsOk() {
		t.Errorf("Expected IsOk to be false for Err result with struct")
	}
	if r2.Value != (S{}) {
		t.Errorf("Expected Value to be zero struct, got %v", r2.Value)
	}
}
