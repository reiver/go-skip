package skip

import (
	"io"

	"sourcecode.social/reiver/go-erorr"
)

// OneExpectedRune skips one expected rune.
// It returns an error if the rune is not what was expected.
func OneExpectedRune(runescanner io.RuneScanner, expected rune) error {
	if nil == runescanner {
		return errNilRuneScanner
	}

	r, size, err := runescanner.ReadRune()
	switch {
	case 0 < size:
		if r != expected {
			return erorr.Errorf("skip:the actual read rune is not what was expected — expected %q (%U), actual %q (%U)", expected, expected, r, r)
		}
		return nil
	case io.EOF == err:
		return io.EOF
	case nil != err:
		return erorr.Errorf("skip: problem skipping one rune: %w", err)
	default:
		return errNoErrorNoRune
	}
}
