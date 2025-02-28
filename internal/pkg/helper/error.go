package helper

import "errors"

var (
	ErrorInvalidPassword                = errors.New("invalid password")
	ErrorUserAlreadyLogout              = errors.New("user already logout")
	ErrorMovieDurationViewershipInvalid = errors.New("invalid movie duration, watching duration more than movie duration")
)
