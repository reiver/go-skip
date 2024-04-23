package skip

import (
	"io"

	"sourcecode.social/reiver/go-erorr"
)

// OneRune skips one rune.
func OneRune(runescanner io.RuneScanner) error {
	if nil == runescanner {
		return errNilRuneScanner
	}

	_, size, err := runescanner.ReadRune()
	switch {
	case 0 < size:
		return nil
	case io.EOF == err:
		return io.EOF
	case nil != err:
		return erorr.Errorf("skip: problem skipping one rune: %w", err)
	default:
		return errNoErrorNoRune
	}
}
