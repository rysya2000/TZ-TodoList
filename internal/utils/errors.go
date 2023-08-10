package utils

import "errors"

var (
	ErrInternalError = errors.New("internal error")
	ErrLenOfTitle = errors.New("title must be filled in and within 200 characters")
	ErrActiveAt = errors.New("wrong format of 'activeAt'. example: 2023-08-04")
	ErrNotFound = errors.New("task by %v ID is not found")
	ErrIDIsNotNum = errors.New("ID is not number") 
)