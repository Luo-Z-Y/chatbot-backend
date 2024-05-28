package internalerror

import "errors"

type ChatNotFoundError struct{}

func (e ChatNotFoundError) Error() string {
	return "Chat not found"
}

func IsChatNotFoundError(err error) bool {
	return errors.Is(err, ChatNotFoundError{})
}

type RequestQueryNotFoundError struct{}

func (e RequestQueryNotFoundError) Error() string {
	return "Request query not found"
}

func IsRequestQueryNotFoundError(err error) bool {
	return errors.Is(err, RequestQueryNotFoundError{})
}
