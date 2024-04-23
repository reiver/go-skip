package skip

import (
	"io"

	"sourcecode.social/reiver/go-erorr"
)

// AnyRunes skips zero, one, or more leading runes that are in the cutset.
func AnyRunes(runescanner io.RuneScanner, cutset ...rune) error {
	if nil == runescanner {
		return errNilRuneScanner
	}

	if len(cutset) <= 0 {
		return nil
	}

	loop: for {
		r, size, err := runescanner.ReadRune()
		switch {
		case 0 < size:
			for _, cut := range cutset {
				if cut == r {
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
}
