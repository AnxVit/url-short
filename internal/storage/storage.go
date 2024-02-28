package storage

import "errors"

var (
	ErrURLNotFOund = errors.New("url not found")
	ErrURLExists   = errors.New("url exists")
)
