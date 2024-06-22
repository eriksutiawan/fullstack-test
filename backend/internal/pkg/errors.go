package pkg

import "errors"

var (
	InternalServerError = errors.New("InternalServerError")
	BadRequestError     = errors.New("BadRequestError")
)
