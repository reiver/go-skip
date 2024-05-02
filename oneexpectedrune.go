package skip

import (
	"io"
)

// OneExpectedRune skips one expected rune.
// It returns an error if the rune is not what was expected.
func OneExpectedRune(runescanner io.RuneScanner, expected rune) error {
	return OneExpectedRunes(runescanner, expected)
}
