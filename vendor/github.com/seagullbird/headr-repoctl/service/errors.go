package service

import (
	"github.com/go-errors/errors"
	"fmt"
)

type ErrPathNotExist error
type ErrUnexpected error

func MakeErrPathNotExist(path string) ErrPathNotExist {
	return errors.New(fmt.Sprintf("Path not exist: %s", path))
}

func MakeErrUnexpected(err error) ErrUnexpected {
	return errors.New(fmt.Sprintf("Unexpected error: %v", err))
}
