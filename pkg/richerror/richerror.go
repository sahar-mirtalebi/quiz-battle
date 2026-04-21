package richerror

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

func New(operation Op) RichError {
	return RichError{
		operation: operation,
	}
}

func (r RichError) WithKind(Kind Kind) RichError {
	return RichError{
		kind: Kind,
	}
}

func (r RichError) WithMessage(message string) RichError {
	return RichError{
		message: message,
	}
}

func (r RichError) WithError(error error) RichError {
	return RichError{
		wrappedError: error,
	}
}

func (r RichError) Withmeta(meta map[string]interface{}) RichError {
	return RichError{
		meta: meta,
	}
}

func (r RichError) Error() string {
	return r.message
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
		return r.wrappedError.Error()
	}

	return re.Message()
}
