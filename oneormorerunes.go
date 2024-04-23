package skip

import (
	"io"

	"sourcecode.social/reiver/go-erorr"
)

// OneOrMoreRunes skips one or more leading runes that are in the cutset.
// It will return the error skin.ErrRuneNotFound if there is not at least one rune read that is in the cutset.
func OneOrMoreRunes(runescanner io.RuneScanner, cutset ...rune) error {
	if nil == runescanner {
		return errNilRuneScanner
	}

	if len(cutset) <= 0 {
		return nil
	}

	var skipped bool

	loop: for {
		r, size, err := runescanner.ReadRune()
		switch {
		case 0 < size:
			for _, cut := range cutset {
				if cut == r {
					skipped = true
					continue loop
				}
			}
		case io.EOF == err:
			return io.EOF
		case nil != err:
			return erorr.Errorf("skip: problem reading rune: %w", err)
		default:
			return errNoErrorNoRune
		}

		{
			err := runescanner.UnreadRune()
			if nil != err {
				return err
			}

			return nil
		}
	}

	if !skipped {
		return ErrRuneNotFound
	}

	return nil
}
