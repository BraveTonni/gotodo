package core_errors

import "errors"

var (
	NotFound = errors.New("Not Found")
	InvalidArgument = errors.New("Invalid Argument")
	Conflict = errors.New("Conflict")
)