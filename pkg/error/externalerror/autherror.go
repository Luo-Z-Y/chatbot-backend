package externalerror

import "errors"

type AuthRequiredError struct{}

func (AuthRequiredError) Error() string {
	return "authentication required"
}

func IsAuthRequiredError(err error) bool {
	return errors.Is(err, AuthRequiredError{})
}
