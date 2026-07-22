// Package richerror adds application metadata and classification to errors.
package richerror

import "runtime"

type Kind int
type Op string

const (
	KindInvalid Kind = iota + 1
	KindForbiden
	KindNotFound
	KindUnexpected
)

type RichError struct {
	operation    Op
	wrappedError error
	message      string
	kind         Kind
	meta         map[string]interface{}
}

func New() RichError {

	pc, _, _, ok := runtime.Caller(1)
	var op Op

	if ok {
		fn := runtime.FuncForPC(pc)
		if fn != nil {
			op = Op(fn.Name())
		}
	}
	return RichError{
		operation: op,
	}

}

func (r RichError) WithKind(Kind Kind) RichError {
	r.kind = Kind

	return r
}

func (r RichError) WithMessage(message string) RichError {
	r.message = message

	return r
}

func (r RichError) WithError(error error) RichError {
	r.wrappedError = error

	return r
}

func (r RichError) WithMeta(meta map[string]interface{}) RichError {
	r.meta = meta

	return r
}

func (r RichError) Error() string {
	return r.Message()
}

func (r RichError) Unwrap() error {
	return r.wrappedError
}

func (r RichError) Kind() Kind {
	if r.kind != 0 {
		return r.kind
	}

	re, ok := r.wrappedError.(RichError)
	if !ok {
		return 0
	}

	return re.Kind()
}

func (r RichError) Message() string {
	if r.message != "" {
		return r.message
	}

	re, ok := r.wrappedError.(RichError)
	if !ok {
		if r.wrappedError == nil {
			return ""
		}
		return r.wrappedError.Error()
	}

	return re.Message()
}
