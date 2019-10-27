package database

import (
	"time"
)

// Parameters
type Parameters struct {
	Created time.Time
	K       int
	M       int
	N       int
}

// Parameters
type List struct {
	Created time.Time
	IsWhite bool
	Network string
}

type Database interface {
	GetParametrs() (*Parameters, error)
	UpdateParametrs(k, m, n int) error
	GetAddressList() (*[]List, error)
	AddAddress(network string, isWhite bool) error
	RemoveAddress(network string, isWhite bool) error
}
