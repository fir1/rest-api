package erru

type ErrArgument struct {
	Wrapped error
}

func (e ErrArgument) Error() string {
	return "invalid argument"
}

func (e ErrArgument) Unwrap() error {
	return e.Wrapped
}
