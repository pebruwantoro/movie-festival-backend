package helper

import "errors"

var (
	ErrorInvalidPassword   = errors.New("invalid password")
	ErrorUserAlreadyLogout = errors.New("user already logout")
)
