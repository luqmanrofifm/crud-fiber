package errs

type (
	BadRequestError struct {
		Err string
	}

	ResourceNotFoundError struct {
		Err string
	}

	UnauthorizedError struct {
		Err string
	}
)

func (e *BadRequestError) Error() string {
	return e.Err
}

func (e *ResourceNotFoundError) Error() string {
	return e.Err
}

func (e *UnauthorizedError) Error() string {
	return e.Err
}
