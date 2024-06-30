package models

import "errors"

var ErrNoRecord = errors.New("no matching record found")
var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrDuplicateEmail = errors.New("duplicate email")
