package stream

import "errors"

var (
	ErrNotFound       = errors.New("not found")
	ErrArrayTypeError = errors.New("array type is not Slice and Array")
)
