/*
Package errors implements functions to manipulate errors.
*/
package errors

import (
	"fmt"

	"google.golang.org/grpc/codes"
)

// Error defines an error with an associated code value.
type Error struct {
	code codes.Code
	desc string
}

// Error implements the builtin error interface.
func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.code, e.desc)
}

// Code returns the error code for `err`. Otherwise, it returns codes.Unknown.
func Code(err error) codes.Code {
	if err == nil {
		return codes.OK
	}
	if e, ok := err.(*Error); ok {
		return e.code
	}
	return codes.Unknown
}

// ErrorDesc returns the error description of `err`. Otherwise, it returns
// err.Error(), or an empty string when `err` is nil.
func ErrorDesc(err error) string {
	if err == nil {
		return ""
	}
	if e, ok := err.(*Error); ok {
		return e.desc
	}
	return err.Error()
}

// Errorf returns an error containing an error code and a description.
// Errorf returns nil if `c` is OK.
func Errorf(c codes.Code, format string, a ...interface{}) error {
	if c == codes.OK {
		return nil
	}
	return &Error{
		code: c,
		desc: fmt.Sprintf(format, a...),
	}
}
