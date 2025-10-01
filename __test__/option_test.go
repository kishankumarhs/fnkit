package fnkit

import (
	"testing"
)

func TestOptionSome(t *testing.T) {
	o := Some(42)
	if !o.IsSome() {
		t.Errorf("Expected IsSome to be true for Some(42)")
	}
	if o.IsNone() {
		t.Errorf("Expected IsNone to be false for Some(42)")
	}
	if o.Unwrap() != 42 {
		t.Errorf("Expected Unwrap to return 42, got %v", o.Unwrap())
	}
	if o.UnwrapOr(99) != 42 {
		t.Errorf("Expected UnwrapOr to return 42, got %v", o.UnwrapOr(99))
	}
}

func TestOptionNone(t *testing.T) {
	o := None[int]()
	if o.IsSome() {
		t.Errorf("Expected IsSome to be false for None[int]()")
	}
	if !o.IsNone() {
		t.Errorf("Expected IsNone to be true for None[int]()")
	}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected Unwrap to panic on None")
		}
	}()
	_ = o.Unwrap() // should panic
}

func TestOptionUnwrapOr(t *testing.T) {
	o := None[string]()
	if o.UnwrapOr("fallback") != "fallback" {
		t.Errorf("Expected UnwrapOr to return fallback, got %v", o.UnwrapOr("fallback"))
	}
	o2 := Some("hello")
	if o2.UnwrapOr("fallback") != "hello" {
		t.Errorf("Expected UnwrapOr to return hello, got %v", o2.UnwrapOr("fallback"))
	}
}

func TestOptionWithStruct(t *testing.T) {
	type S struct{ X int }
	o := Some(S{X: 7})
	if !o.IsSome() {
		t.Errorf("Expected IsSome to be true for Some(struct)")
	}
	if o.Unwrap().X != 7 {
		t.Errorf("Expected Unwrap().X to be 7, got %v", o.Unwrap().X)
	}
	n := None[S]()
	if n.IsSome() {
		t.Errorf("Expected IsSome to be false for None[S]()")
	}
	if n.UnwrapOr(S{X: 99}).X != 99 {
		t.Errorf("Expected UnwrapOr to return fallback struct, got %v", n.UnwrapOr(S{X: 99}))
	}
}
