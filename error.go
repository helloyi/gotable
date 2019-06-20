package table

import (
	"reflect"
)

type (
	// ErrNumOverflow ...
	ErrNumOverflow struct {
		Method string
		Kind   reflect.Kind
	}

	// ErrUnsupportedKind ...
	ErrUnsupportedKind struct {
		Method string
		Kind   interface{}
	}

	// ErrCannotBeNil ...
	ErrCannotBeNil struct {
		Method string
	}

	ErrNotExist struct {
		Method string
		Thing  string
	}
)

func (e *ErrUnsupportedKind) Error() string {
	rkind, ok := e.Kind.(reflect.Kind)
	if ok && rkind == 0 {
		return "table: call of " + e.Method + " on zero value"
	}

	var kind string
	if ok {
		kind = rkind.String()
	} else {
		kind, _ = e.Kind.(string)
	}

	return "table: call of " + e.Method + " on " + kind + " value"
}

func (e *ErrNumOverflow) Error() string {
	return "table: call of " + e.Method + " overflows " + e.Kind.String()
}

func (e *ErrCannotBeNil) Error() string {
	return "table: call of " + e.Method + " on nil value"
}

func (e *ErrNotExist) Error() string {
	return "table: call of " + e.Method + " not exist of " + e.Thing
}
