package errors

type AppError struct {
	Code int

	Key string

	Err error
}

func (e *AppError) Error() string {

	return e.Key
}

func (e *AppError) Unwrap() error {

	return e.Err
}
