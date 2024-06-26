package skip

import (
	"sourcecode.social/reiver/go-erorr"
)

const (
	errNilRuneScanner = erorr.Error("skip: nil rune scanner")
	errNoErrorNoRune  = erorr.Error("skip: problem reading rune: did not receive an error but also did not receive a rune")
)

const (
	ErrRuneNotFound = erorr.Error("rune not found")
)
