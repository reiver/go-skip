package skip

import (
	"io"
)

// AnyRunes skips zero, one, or more leading runes that are in the cutset.
func AnyRunes(runescanner io.RuneScanner, cutset ...rune) error {
	err := OneOrMoreRunes(runescanner, cutset...)
	if ErrRuneNotFound == err {
		return nil
	}

	return err
}
