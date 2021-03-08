package errors

import "errors"

func NewInvalid(errorMsg string) error {
	return errors.New(errorMsg)
}
