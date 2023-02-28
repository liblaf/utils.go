package errors

import "errors"

func Unreachable() error {
	return errors.New("unreachable")
}

func NotImplemented() error {
	return errors.New("not implemented")
}
